package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduledAlertAt(duration time.Duration, amount int)
}

type CLI struct {
	PlayerStore PlayerStore
	In          *bufio.Scanner
	Alerter     BlindAlerter
}

func NewCli(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
		Alerter:     alerter,
	}
}

func (cli *CLI) PlayPoker() {
	cli.Alerter.ScheduledAlertAt(5*time.Second, 100)
	userInput := cli.readLine()
	cli.PlayerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}
