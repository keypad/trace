package core

import (
	"fmt"
	"strings"
	"time"
)

func dnstable(row dnsrow) string {
	lines := []string{"target  state  count  ips  note"}
	lines = append(lines, fmt.Sprintf("%s  %s  %d  %s  %s", row.target, row.state, row.count, clean(row.ips), clean(row.note)))
	return strings.Join(lines, "\n") + "\n"
}

func tlstable(row tlsrow) string {
	lines := []string{"target  state  version  expires  days  issuer  note"}
	lines = append(lines, fmt.Sprintf("%s  %s  %s  %s  %d  %s  %s", row.target, row.state, clean(row.version), clean(row.expires), row.days, clean(row.issuer), clean(row.note)))
	return strings.Join(lines, "\n") + "\n"
}

func httptable(row httprow) string {
	lines := []string{"target  state  code  latency  size  note"}
	lines = append(lines, fmt.Sprintf("%s  %s  %d  %s  %s  %s", row.target, row.state, row.code, latency(row.latency), size(row.size), clean(row.note)))
	return strings.Join(lines, "\n") + "\n"
}

func clean(value string) string {
	if strings.TrimSpace(value) == "" {
		return "-"
	}
	return strings.ReplaceAll(value, " ", ",")
}

func size(value int64) string {
	if value <= 0 {
		return "-"
	}
	return fmt.Sprintf("%d", value)
}

func latency(value time.Duration) string {
	if value <= 0 {
		return "-"
	}
	return value.Round(time.Millisecond).String()
}
