package select_race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func urlComparator(t *testing.T, urlA, urlB string) {
	if urlA != urlB {
		t.Errorf("got %q, want %q", urlA, urlB)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRace(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		slowURL := "http://www.facebook.com"
		fastURL := "http://www.quii.dev"

		want := fastURL
		got, time, err := Racer(slowURL, fastURL)
		t.Logf("time taken: %v", time)
		urlComparator(t, got, want)
		if err != nil {
			t.Fatalf("got an error %v", err)
		}

	})

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, time, err := Racer(slowURL, fastURL)
		t.Logf("time taken: %v", time)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, time, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		t.Logf("time taken: %v", time)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
