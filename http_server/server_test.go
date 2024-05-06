package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type StubPlayersStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayersStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}

	server := mustMakePlayerServer(t, &store)

	t.Run("return Pepper's scores", func(t *testing.T) {
		req := NewGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertStatusCode(t, res.Code, http.StatusOK)
		AssertResponseBody(t, res.Body.String(), "20")
	})
	t.Run("return Floyd's scores", func(t *testing.T) {
		req := NewGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertStatusCode(t, res.Code, http.StatusOK)
		AssertResponseBody(t, res.Body.String(), "10")
	})

	t.Run("return 404 error on missing players", func(t *testing.T) {
		req := NewGetScoreRequest("Apollo")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertStatusCode(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayersStore{
		map[string]int{},
		nil,
		nil,
	}
	server := mustMakePlayerServer(t, &store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		req := newPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertStatusCode(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayersStore{}
	server := mustMakePlayerServer(t, &store)

	t.Run("it returns status 200 on /league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var got []Player

		err := json.NewDecoder(res.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", res.Body, err)
		}

		AssertStatusCode(t, res.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayersStore{nil, nil, wantedLeague}
		server := mustMakePlayerServer(t, &store)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := getLeagueFromResponse(t, res.Body)
		AssertStatusCode(t, res.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)

		AssertContentType(t, res, jsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server, _ := NewPlayerServer(&StubPlayersStore{})

		req := newGameRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertStatusCode(t, res.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket it is a winner of a game", func(t *testing.T) {
		store := &StubPlayersStore{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, store))
		defer server.Close()

		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer ws.Close()

		writeWSMessage(t, ws, winner)
		time.Sleep(10 * time.Millisecond)
		AssertPlayerWin(t, store, winner)
	})
}

func (s *StubPlayersStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayersStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayersStore) GetLeague() League {
	return s.league
}

func mustMakePlayerServer(t *testing.T, store PlayerStore) *PlayerServer {
	server, err := NewPlayerServer(store)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func AssertPlayerWin(t testing.TB, store *StubPlayersStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}
