package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type CLI struct {
	PlayerStore PlayerStore
	In          *bufio.Scanner
	Out         io.Writer
	game        Game
}

func NewCli(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		In:   bufio.NewScanner(in),
		Out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.Out, PlayerPrompt)

	// numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(cli.readLine())

	if err != nil {
		fmt.Fprint(cli.Out, "you're so silly")
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
