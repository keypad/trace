```bash
> what is this?

  dns + tls + http target inspector utility.
  local cli + wttr-style plain-text http output.
  built for terminal and curl workflows.

> features?

  ✓ dns lookup with timeout control
  ✓ tls certificate and protocol inspection
  ✓ http status + latency checks
  ✓ plain-text server mode for curl usage
  ✓ no credentials, no database, no external accounts

> usage?

  trace dns <host> [timeoutms]
  trace tls <host[:port]> [timeoutms]
  trace http <url> [timeoutms]
  trace serve [port] [timeoutms]

> examples?

  go run ./src dns localhost
  go run ./src tls example.com:443
  go run ./src http https://example.com
  go run ./src serve 4195
  curl "http://127.0.0.1:4195/http?url=https://example.com"

> stack?

  go 1.26 stdlib

> run?

  go run ./src dns localhost
  go run ./src serve 4195

> test?

  go test ./...

> proof?

  $ go test ./...
  ?    github.com/keypad/trace/src       [no test files]
  ?    github.com/keypad/trace/src/core  [no test files]
  ok   github.com/keypad/trace/test      0.229s

  $ go run ./src dns localhost 1500
  target  state  count  ips  note
  localhost  up  2  127.0.0.1,::1  -

  $ curl "http://127.0.0.1:4195/http?url=https://example.com"
  target  state  code  latency  size  note
  https://example.com  up  200  70ms  -  -

> links?

  https://github.com/keypad/trace
```
