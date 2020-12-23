package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pippen's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pippen", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns Jordan's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Jordan", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "45"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
