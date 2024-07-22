package tester

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

const (
	BUFF_SIZE = (32 * 1024)
)

type Tester struct {
	Wg      *sync.WaitGroup
	Mu      *sync.Mutex
	LimitCh chan bool
	Config  config
}

type ipInfo struct {
	Ip    string
	Bytes int
}

func Init(confPath string) (*Tester, error) {
    if Limit <= 0 {
        Limit = DEFAULT_LIMIT
    }

	t := &Tester{
		Wg:      &sync.WaitGroup{},
		Mu:      &sync.Mutex{},
		LimitCh: make(chan bool, Limit),
	}

	cdata, err := os.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(cdata, &t.Config)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func createClient(ip string) *http.Client {
	dialer := &net.Dialer{
		// Lookup timeout
		Timeout: LOOKUP_TIMEOUT,
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial(network, net.JoinHostPort(ip, PORT))
			},
		},
	}

	return &http.Client{
		// Request timeout
		Timeout: REQUEST_TIMEOUT,
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
		},
	}
}

func (t *Tester) ipIsOk(info *ipInfo) bool {
	c := createClient(info.Ip)

	req, err := http.NewRequestWithContext(context.Background(), "GET", t.Config.Url, nil)
	if err != nil {
		return false
	}

	resp, err := c.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	buff := make([]byte, BUFF_SIZE)
	reader := bufio.NewReaderSize(resp.Body, BUFF_SIZE)

	if resp.StatusCode == http.StatusOK {
		for {
			read, err := reader.Read(buff)
			info.Bytes += read
			if err != nil {
				break
			}
		}
		return true
	}

	return false
}

func (t *Tester) TestIPs() map[string]int {
	okIPs := make(map[string]int)

	for _, ip := range t.Config.Ips {
		t.Wg.Add(1)
		go func(ip string) {
			defer t.Wg.Done()
			t.LimitCh <- false

			info := ipInfo{ip, 0}
			if t.ipIsOk(&info) {
				log.Printf("[OK] %s\n", ip)
				okIPs[ip] = info.Bytes / 1024
			} else {
				log.Printf("[FAIL] %s\n", ip)
			}

            <-t.LimitCh
		}(ip)
	}
	t.Wg.Wait()

	return okIPs
}
