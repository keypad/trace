package core

import (
	"context"
	"net"
	"sort"
	"strings"
	"time"
)

func dnscheck(target string, timeout time.Duration) dnsrow {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	addresses, err := net.DefaultResolver.LookupIPAddr(ctx, target)
	if err != nil {
		return dnsrow{target: target, state: "down", count: 0, ips: "", note: err.Error()}
	}

	values := make([]string, 0, len(addresses))
	for _, address := range addresses {
		values = append(values, address.IP.String())
	}
	sort.Strings(values)

	return dnsrow{target: target, state: "up", count: len(values), ips: strings.Join(values, ","), note: ""}
}
