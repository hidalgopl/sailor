package auth

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)


// NewTestServer starts new *httptest.Server with a given host and port
func NewTestServer(host, port string, handler http.Handler) *httptest.Server {
	hostPort := fmt.Sprintf("%s:%s", host, port)
	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatal(err)
	}
	ts := httptest.NewUnstartedServer(handler)
	ts.Listener = l
	ts.Start()
	return ts
}


func TestDoAuth(t *testing.T) {
	// Start a local HTTP server
	newMockWDServer := func() *httptest.Server {
		r := http.NewServeMux()
		r.HandleFunc("/wd/hub/session", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			body := []byte(`{"sessionId":"` +  `"}`)
			_, err := w.Write(body)
			if err != nil {
				log.Fatal(err)
			}
		}))
		r.HandleFunc("/wd/hub/session/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body := []byte(`{"src":"` + `"}`)
			_, err := w.Write(body)
			if err != nil {
				log.Fatal(err)
			}
		}))
		server := NewTestServer("0.0.0.0", "8072", r)
		return server
	}
	mockServer := newMockWDServer()
	// Close the server when test finishes
	defer mockServer.Close()
}
