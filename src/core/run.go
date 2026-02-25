package core

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func Run(args []string, out io.Writer, errout io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(errout, usage())
		return 1
	}

	timeout := 3 * time.Second

	switch args[0] {
	case "dns":
		if len(args) < 2 {
			fmt.Fprint(errout, usage())
			return 1
		}
		if len(args) > 2 {
			value, err := parsems(args[2])
			if err != nil {
				fmt.Fprintln(errout, err.Error())
				return 1
			}
			timeout = value
		}
		fmt.Fprint(out, dnstable(dnscheck(args[1], timeout)))
		return 0
	case "tls":
		if len(args) < 2 {
			fmt.Fprint(errout, usage())
			return 1
		}
		if len(args) > 2 {
			value, err := parsems(args[2])
			if err != nil {
				fmt.Fprintln(errout, err.Error())
				return 1
			}
			timeout = value
		}
		fmt.Fprint(out, tlstable(tlscheck(args[1], timeout)))
		return 0
	case "http":
		if len(args) < 2 {
			fmt.Fprint(errout, usage())
			return 1
		}
		if len(args) > 2 {
			value, err := parsems(args[2])
			if err != nil {
				fmt.Fprintln(errout, err.Error())
				return 1
			}
			timeout = value
		}
		fmt.Fprint(out, httptable(httpcheck(args[1], timeout)))
		return 0
	case "serve":
		port := "4175"
		if len(args) > 1 {
			port = args[1]
		}
		if len(args) > 2 {
			value, err := parsems(args[2])
			if err != nil {
				fmt.Fprintln(errout, err.Error())
				return 1
			}
			timeout = value
		}
		if err := serve(port, timeout); err != nil {
			fmt.Fprintln(errout, err.Error())
			return 1
		}
		return 0
	default:
		fmt.Fprint(errout, usage())
		return 1
	}
}

func parsems(value string) (time.Duration, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, fmt.Errorf("timeoutms must be a positive integer")
	}
	number, err := strconv.Atoi(value)
	if err != nil || number <= 0 {
		return 0, fmt.Errorf("timeoutms must be a positive integer")
	}
	return time.Duration(number) * time.Millisecond, nil
}

func usage() string {
	return "trace dns <host> [timeoutms]\ntrace tls <host[:port]> [timeoutms]\ntrace http <url> [timeoutms]\ntrace serve [port] [timeoutms]\n"
}
