package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// // This passes the tests sometimes & fails sometimes because depends on flakey external servers
// func TestRacer(t *testing.T) {
// 	slowURL := "http://www.facebook.com"
// 	fastURL := "http://www.quii.dev"

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// this test uses a internally run mock server to simulate a certain slow behaviour
// ie. a slow server, and shows that indeed the function is returning the faster server


func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

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