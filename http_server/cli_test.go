package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayersStore{}

		cli := NewCli(playerStore, in)
		cli.PlayPoker()

		if len(playerStore.winCalls) != 1 {
			t.Fatal("expected a win call but didn't get any")
		}

		AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayersStore{}

		cli := NewCli(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Cleo")
	})

}
