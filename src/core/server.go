package core

import (
	"net/http"
	"time"
)

func Handler(timeout time.Duration) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/dns", func(writer http.ResponseWriter, request *http.Request) {
		target := request.URL.Query().Get("host")
		if target == "" {
			http.Error(writer, "missing host", http.StatusBadRequest)
			return
		}
		ms := timeout
		if value := request.URL.Query().Get("timeoutms"); value != "" {
			parsed, err := parsems(value)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			ms = parsed
		}
		writer.Header().Set("content-type", "text/plain; charset=utf-8")
		_, _ = writer.Write([]byte(dnstable(dnscheck(target, ms))))
	})
	mux.HandleFunc("/tls", func(writer http.ResponseWriter, request *http.Request) {
		target := request.URL.Query().Get("host")
		if target == "" {
			http.Error(writer, "missing host", http.StatusBadRequest)
			return
		}
		ms := timeout
		if value := request.URL.Query().Get("timeoutms"); value != "" {
			parsed, err := parsems(value)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			ms = parsed
		}
		writer.Header().Set("content-type", "text/plain; charset=utf-8")
		_, _ = writer.Write([]byte(tlstable(tlscheck(target, ms))))
	})
	mux.HandleFunc("/http", func(writer http.ResponseWriter, request *http.Request) {
		target := request.URL.Query().Get("url")
		if target == "" {
			http.Error(writer, "missing url", http.StatusBadRequest)
			return
		}
		ms := timeout
		if value := request.URL.Query().Get("timeoutms"); value != "" {
			parsed, err := parsems(value)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			ms = parsed
		}
		writer.Header().Set("content-type", "text/plain; charset=utf-8")
		_, _ = writer.Write([]byte(httptable(httpcheck(target, ms))))
	})
	return mux
}

func serve(port string, timeout time.Duration) error {
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           Handler(timeout),
		ReadHeaderTimeout: 5 * time.Second,
	}
	return server.ListenAndServe()
}
