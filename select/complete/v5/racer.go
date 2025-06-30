package racer5

import (
	"fmt"
	"net/http"
	"time"
)

// V5 用 timeout
func Racer(a, b string) (string, error) {
	select {
	case <-ping(a): // block住, 看誰先回傳解block
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second): // chan, 最後會會傳當下時間time.Time
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch) // 直接回一個 zero value
	}()

	return ch
}
