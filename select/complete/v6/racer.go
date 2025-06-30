package racer6

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// V5 用 timeout
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer 還是想不太到要怎麼直覺得想到要拔出來
// 1. 看到 magic number / magic decision ── 把它升格成「策略」
// 2. 測試出現 非必要的等待 / 隨機性 / 外部依賴
// 3. 需求文件或同事口頭反覆提到「要能調」
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a): // block住, 看誰先回傳解block
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // chan, 最後會會傳當下時間time.Time
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
