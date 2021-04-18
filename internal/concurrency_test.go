package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// https://dev.to/quii/learn-go-by-writing-tests-synchronising-asynchronous-processes-1bo7

func TestRacer(t *testing.T) {
	t.Run("returns fastest url to load", func(t *testing.T) {
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
		actURL, err := Racer(slowURL, fastURL, time.Duration(20*time.Millisecond))

		// then
		require.NoError(t, err)
		require.Equal(t, fastURL, actURL)
	})
	t.Run("returns an error if url load times out", func(t *testing.T) {
		//	given
		slowServer := httpTestServer(20 * time.Millisecond)
		fastServer := httpTestServer(10 * time.Millisecond)

		slowServer.Start()
		defer slowServer.Close()
		fastServer.Start()
		defer fastServer.Close()
		//	when
		url, err := Racer(slowServer.URL, fastServer.URL, time.Duration(1*time.Millisecond))
		fmt.Printf("slowUrl: %s, fastUrl: %s, fasterUrl: %s\n",
			slowServer.URL, fastServer.URL, url)
		//	then
		require.Error(t, err)
	})
}

func httpTestServer(delay time.Duration) httptest.Server {
	return *httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
