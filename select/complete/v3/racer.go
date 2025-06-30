package racer3

import (
	"net/http"
	"time"
)

// V2 改用httptest server 模擬外部
func Racer(a, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}
