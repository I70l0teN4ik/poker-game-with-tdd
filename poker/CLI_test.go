package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"tdd/poker"
	"testing"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}


func Test_CLI(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		cases := []struct{
			numOfPlayers int
			winner string
		}{
			{4, "Artem"},
			{6, "Rick"},
			{9, "Morty"},
		}
		for _, tt := range cases {
			t.Run(fmt.Sprintf("%s wins a game with %d players", tt.winner, tt.numOfPlayers), func(t *testing.T) {
				stdout := &bytes.Buffer{}
				in := userInput(tt.numOfPlayers, tt.winner + " wins")
				game := &GameSpy{}

				cli := poker.NewCLI(in, stdout, game)
				cli.PlayPoker()

				assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
				assertGameStartedWith(t, game, tt.numOfPlayers)
				assertFinishCalledWith(t, game, tt.winner)
			})
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userInput("four")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt+poker.BadPlayerInputErrMsg)
	})
}


func userInput(inputs ...interface{}) *strings.Reader {
	s := ""
	for _, input := range inputs {
		s += fmt.Sprintf("%v\n", input)
	}
	return strings.NewReader(s)
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	if game.FinishedWith != winner {
		t.Errorf("wanted %s to be a winner but got %s", winner, game.FinishedWith)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, count int) {
	t.Helper()
	if game.StartedWith != count {
		t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
