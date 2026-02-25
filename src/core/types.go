package core

import "time"

type dnsrow struct {
	target string
	state  string
	count  int
	ips    string
	note   string
}

type tlsrow struct {
	target  string
	state   string
	version string
	expires string
	days    int
	issuer  string
	note    string
}

type httprow struct {
	target  string
	state   string
	code    int
	latency time.Duration
	size    int64
	note    string
}
