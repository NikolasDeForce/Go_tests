package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	PlayerStore PlayerStore
	In          io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.In)
	reader.Scan()
	cli.PlayerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
