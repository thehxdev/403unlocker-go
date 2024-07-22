package tester

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

type Tester struct {
	Wg      *sync.WaitGroup
	Mu      *sync.Mutex
	LimitCh chan bool
	Config  config
}

func Init(confPath string) (*Tester, error) {
	t := &Tester{
		Wg:      &sync.WaitGroup{},
		Mu:      &sync.Mutex{},
		LimitCh: make(chan bool, LIMIT),
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

func (t *Tester) ipIsOk(ip string) bool {
	c := createClient(ip)

	req, err := http.NewRequestWithContext(context.Background(), "GET", t.Config.Url, nil)
	if err != nil {
		return false
	}

	resp, err := c.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func (t *Tester) TestIPs() []string {
	okIPs := make([]string, 0)

	for _, ip := range t.Config.Ips {
		t.Wg.Add(1)
		go func(ip string) {
			defer t.Wg.Done()
			t.LimitCh <- t.ipIsOk(ip)

			if <-t.LimitCh {
				log.Printf("[OK] %s\n", ip)
				t.Mu.Lock()
				okIPs = append(okIPs, ip)
				t.Mu.Unlock()
			}
		}(ip)
	}
	t.Wg.Wait()

	return okIPs
}
