// Package v3 增加time.sleep 等 go routine都好了
package v3

import "time"

// WebsiteChecker 負責確認 url 的狀況
// 利用DI方便mock
type WebsiteChecker func(string) bool

// CheckWebsites 確認 url 的狀態
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	time.Sleep(2 * time.Second)

	return results
}
