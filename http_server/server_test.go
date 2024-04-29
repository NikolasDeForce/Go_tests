package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayersStore struct {
	scores map[string]int
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayersStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("return Pepper's scores", func(t *testing.T) {
		req := NewGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "20")
	})
	t.Run("return Floyd's scores", func(t *testing.T) {
		req := NewGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "10")
	})

	t.Run("return 404 error on missing players", func(t *testing.T) {
		req := NewGetScoreRequest("Apollo")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusNotFound)
	})
}

func (s *StubPlayersStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
