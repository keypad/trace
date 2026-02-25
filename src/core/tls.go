package core

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"
)

func tlscheck(target string, timeout time.Duration) tlsrow {
	host, port := split(target)
	address := net.JoinHostPort(host, port)

	dialer := net.Dialer{Timeout: timeout}
	connection, err := tls.DialWithDialer(&dialer, "tcp", address, &tls.Config{ServerName: host, MinVersion: tls.VersionTLS12, InsecureSkipVerify: true})
	if err != nil {
		return tlsrow{target: target, state: "down", version: "", expires: "", days: 0, issuer: "", note: err.Error()}
	}
	defer connection.Close()

	state := connection.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		return tlsrow{target: target, state: "down", version: "", expires: "", days: 0, issuer: "", note: "missing certificate"}
	}

	leaf := state.PeerCertificates[0]
	days := int(time.Until(leaf.NotAfter).Hours() / 24)
	issuer := leaf.Issuer.CommonName
	if issuer == "" && len(leaf.Issuer.Organization) > 0 {
		issuer = leaf.Issuer.Organization[0]
	}

	return tlsrow{
		target:  target,
		state:   "up",
		version: version(state.Version),
		expires: leaf.NotAfter.UTC().Format(time.RFC3339),
		days:    days,
		issuer:  issuer,
		note:    "",
	}
}

func split(target string) (string, string) {
	value := strings.TrimSpace(target)
	if value == "" {
		return "", "443"
	}

	if strings.Contains(value, ":") {
		host, port, err := net.SplitHostPort(value)
		if err == nil {
			if port == "" {
				return host, "443"
			}
			return host, port
		}
	}

	return value, "443"
}

func version(value uint16) string {
	switch value {
	case tls.VersionTLS10:
		return "tls1.0"
	case tls.VersionTLS11:
		return "tls1.1"
	case tls.VersionTLS12:
		return "tls1.2"
	case tls.VersionTLS13:
		return "tls1.3"
	default:
		return fmt.Sprintf("0x%x", value)
	}
}
