package poker

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

var (
	dummySpyAlerter   = &SpyBlindAlerter{}
	dummyBlindAlerter = &SpyBlindAlerter{}
	dummyPlayerStore  = &StubPlayersStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}

		cli := NewCli(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris', but got %q", game.FinishedWith)
		}
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		game := &GameSpy{}

		cli := NewCli(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo', but got %q", game.FinishedWith)
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := NewCli(in, stdOut, game)
		cli.PlayPoker()

		gotPrompt := stdOut.String()
		wantPrompt := PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := NewCli(in, stdOut, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})
}

func AssertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduledAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}
