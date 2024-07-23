package tester

import (
	"runtime"
)

const (
	// DNS server port
	PORT = "53"
)

var (
	// Process this number of IPs concurrently
	DEFAULT_LIMIT = max(1, runtime.NumCPU())
	Limit         = DEFAULT_LIMIT
)

type config struct {
	Url            string   `json:"url"`
	Ips            []string `json:"ips"`
	LookupTimeout  int      `json:"lookup_timeout"`
	RequestTimeout int      `json:"request_timeout"`
}
