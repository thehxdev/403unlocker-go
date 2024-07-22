package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thehxdev/403unlocker-go/tester"
)

const VERSION = "0.1.0"

var versionInfo = "403unlocker-go v" + VERSION + "\nhttps://github.com/thehxdev/403unlocker-go"

func main() {
	cpath := flag.String("c", "config.json", "path to config file")
	version := flag.Bool("v", false, "show version info")
	flag.Parse()

	if *version {
		fmt.Println(versionInfo)
		os.Exit(0)
	}

	t, err := tester.Init(*cpath)
	if err != nil {
		log.Fatal(err)
	}

	ips := t.TestIPs()

	fmt.Printf("\nWorking IPs = %+v\n", ips)
}
