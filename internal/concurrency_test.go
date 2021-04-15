package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// https://dev.to/quii/learn-go-by-writing-tests-synchronising-asynchronous-processes-1bo7


func TestRacer(t *testing.T) {
	// given
	slowServer := httpTestServer(20 * time.Millisecond)
	fastServer := httpTestServer(2 * time.Millisecond)

	slowServer.Start()
	defer slowServer.Close()
	fastServer.Start()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	// when
    actURL := Racer(slowURL, fastURL)

	// then
	require.Equal(t, fastURL, actURL)
}


func httpTestServer(delay time.Duration) httptest.Server{
	return *httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}
