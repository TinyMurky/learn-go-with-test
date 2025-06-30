package racer2

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// V2 改用server了
func TestRacer(t *testing.T) {
	slowerServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fasterServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowerServer.URL
	fastURL := fasterServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowerServer.Close()
	fasterServer.Close()
}
