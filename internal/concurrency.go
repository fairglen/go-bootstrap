package internal

import (
	"net/http"
	"time"
)

// Racer returns the fastest url
func Racer(urls ... string) (fasterURL string) {
	var fastestResponseTime time.Duration
	for i, url := range urls {
		currResponseTime := measureResponseTime(url)
		if (currResponseTime < fastestResponseTime || i == 0){
			fasterURL = url
			fastestResponseTime = currResponseTime
		}

	}
	return fasterURL
}


func measureResponseTime(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
