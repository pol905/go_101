package selecter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("when servers take less than 10 seconds", func(t *testing.T) {
		slowServer := makeMockServer(20 * time.Millisecond)
		fastServer := makeMockServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("wasn't expecting an error here")
		}

		if got != want {
			t.Errorf("Got: %q, want: %q", got, want)
		}
	})
	t.Run("when servers take more than 10 seconds to complete", func(t *testing.T) {
		slowServer := makeMockServer(15 * time.Second)
		fastServer := makeMockServer(13 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 2*time.Second)
		if err == nil {
			t.Errorf("expected an error but didn't get any")
		}
	})
}

func makeMockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
