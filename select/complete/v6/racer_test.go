package racer6

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// V5 作refactor
func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {

		slowerServer := makeDelayedServer(5 * time.Millisecond)
		fasterServer := makeDelayedServer(0 * time.Millisecond)

		defer slowerServer.Close()
		defer fasterServer.Close()

		slowURL := slowerServer.URL
		fastURL := fasterServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(1 * time.Microsecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 1*time.Nanosecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
