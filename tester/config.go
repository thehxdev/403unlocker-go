package tester

import (
	"time"
)

const (
	LOOKUP_TIMEOUT  = time.Second * 6
	REQUEST_TIMEOUT = time.Second * 6

	// Process this number of IPs concurrently
	DEFAULT_LIMIT = 2

	// DNS server port
	PORT = "53"
)

var (
    Limit = DEFAULT_LIMIT
)

type config struct {
	Url string   `json:"url"`
	Ips []string `json:"ips"`
}
