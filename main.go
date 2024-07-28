package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/thehxdev/403unlocker-go/tester"
)

const VERSION = "1.1.1"

var versionInfo = "403unlocker-go v" + VERSION + "\nhttps://github.com/thehxdev/403unlocker-go"

func main() {
	cpath := flag.String("c", "config.json", "path to config file")
	version := flag.Bool("v", false, "show version info")
	downloadConfig := flag.Bool("dc", false, "download default config file")
	flag.IntVar(&tester.Limit, "l", tester.DEFAULT_LIMIT, "number of IPs that will be processed concurrently")
	flag.Parse()

	if *version {
		fmt.Println(versionInfo)
		os.Exit(0)
	}

	if *downloadConfig {
		downloadDefaultConfig()
		os.Exit(0)
	}

	t, err := tester.Init(*cpath)
	if err != nil {
		log.Fatal(err)
	}

	ips := t.TestIPs()
	fmt.Printf("\nTested IPs = %+v\n", ips)

	writeToFile(ips)
}

func writeToFile(ips map[string]int) {
	createdTime := time.Now().Format(time.DateTime)
	name := "result-" + createdTime + ".json"

	fp, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	err = json.NewEncoder(fp).Encode(&ips)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote test result to `%s` file\n", name)
}

func downloadDefaultConfig() {
	url := "https://raw.githubusercontent.com/thehxdev/403unlocker-go/main/config.json"
	c := tester.CreateHttpClient("9.9.9.9", 10, 10)
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.Create("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	fp.Write(body)
}
