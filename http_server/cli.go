package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	PlayerStore PlayerStore
	In          *bufio.Scanner
}

func NewCli(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
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
