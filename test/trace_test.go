package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/keypad/trace/src/core"
)

func TestUsage(t *testing.T) {
	out := &bytes.Buffer{}
	errout := &bytes.Buffer{}
	code := core.Run([]string{}, out, errout)
	if code != 1 {
		t.Fatalf("want 1 got %d", code)
	}
	if !strings.Contains(errout.String(), "trace dns") {
		t.Fatalf("missing usage output: %s", errout.String())
	}
}

func TestHttp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "ok")
	}))
	defer server.Close()

	out := &bytes.Buffer{}
	errout := &bytes.Buffer{}
	code := core.Run([]string{"http", server.URL, "1500"}, out, errout)
	if code != 0 {
		t.Fatalf("want 0 got %d err=%s", code, errout.String())
	}
	if !strings.Contains(out.String(), "up") || !strings.Contains(out.String(), "200") {
		t.Fatalf("unexpected output: %s", out.String())
	}
}

func TestTls(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	out := &bytes.Buffer{}
	errout := &bytes.Buffer{}
	code := core.Run([]string{"tls", parsed.Host, "2000"}, out, errout)
	if code != 0 {
		t.Fatalf("want 0 got %d err=%s", code, errout.String())
	}
	if !strings.Contains(out.String(), "tls") || !strings.Contains(out.String(), "up") {
		t.Fatalf("unexpected output: %s", out.String())
	}
}

func TestServerhttp(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusAccepted)
	}))
	defer upstream.Close()

	server := httptest.NewServer(core.Handler(2 * time.Second))
	defer server.Close()

	response, err := http.Get(server.URL + "/http?url=" + url.QueryEscape(upstream.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("want 200 got %d", response.StatusCode)
	}
	text := string(body)
	if !strings.Contains(text, "202") || !strings.Contains(text, "up") {
		t.Fatalf("unexpected output: %s", text)
	}
}
