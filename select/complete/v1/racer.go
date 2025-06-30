package racer1

import (
	"net/http"
	"time"
)

// v1 用 timeSince測試時間
// test 直接用 真的網路資料
func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}
