package racer4

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// V4 作refactor
func TestRacer(t *testing.T) {
	slowerServer := makeDelayedServer(5 * time.Millisecond)
	fasterServer := makeDelayedServer(0 * time.Millisecond)

	defer slowerServer.Close()
	defer fasterServer.Close()

	slowURL := slowerServer.URL
	fastURL := fasterServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
