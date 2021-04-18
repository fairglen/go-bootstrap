package internal

import (
	"fmt"
	"net/http"
	"time"
)

// Racer returns the fastest url
func Racer(urlA, urlB string, timeout time.Duration) (string, error) {
	select {
		case <- ping(urlA):
			return urlA, nil
		case <- ping(urlB):
			return urlB, nil
		case <- time.After(timeout):
			return "", fmt.Errorf("timed out after waiting for %s and %s to load", urlA, urlB)
	}
}

func ping(url string) chan bool{
	ch:= make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

