package racer4

import (
	"net/http"
)

// V4 用Ping
func Racer(a, b string) string {
	select {
	case <-ping(a): // block住, 看誰先回傳解block
		return a
	case <-ping(b):
		return b
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
