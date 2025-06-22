// Package concurrency do concurrency stuff
package concurrency

// WebsiteChecker check if usl is  alive
type WebsiteChecker func(string) bool

type result struct {
	url   string
	alive bool
}

// CheckWebsites check if url is alive
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// https://go.dev/blog/loopvar-preview
	// 1.22 以後 url 就不用顯示傳入 goroutine
	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range len(urls) {
		r := <-resultChannel
		results[r.url] = r.alive
	}

	return results
}
