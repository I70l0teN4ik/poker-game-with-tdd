package poker_test

import (
	"fmt"
	"io"
	"tdd/poker"
	"testing"
	"time"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

func TestGame_Start(t *testing.T) {
	interval := poker.AlertInterval
	t.Run("schedules Alerts on playGame start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(5, io.Discard)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * interval, Amount: 200},
			{At: 20 * interval, Amount: 300},
			{At: 30 * interval, Amount: 400},
			{At: 40 * interval, Amount: 500},
			{At: 50 * interval, Amount: 600},
			{At: 60 * interval, Amount: 800},
			{At: 70 * interval, Amount: 1000},
			{At: 80 * interval, Amount: 2000},
			{At: 90 * interval, Amount: 4000},
			{At: 100 * interval, Amount: 8000},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules Alerts on playGame start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(7, io.Discard)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * interval, Amount: 200},
			{At: 24 * interval, Amount: 300},
			{At: 36 * interval, Amount: 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	tests := []struct {
		playerName string
	}{
		{"Artem"},
		{"Rick"},
	}
	for _, tt := range tests {
		t.Run(tt.playerName+" wins from cli", func(t *testing.T) {
			store := &poker.StubPlayerStore{}
			game := poker.NewGame(dummyBlindAlerter, store)

			game.Finish(tt.playerName)
			poker.AssertPlayerWin(t, store, tt.playerName)
		})
	}
}

func checkSchedulingCases(t *testing.T, cases []poker.ScheduledAlert, alerter *poker.SpyBlindAlerter) {
	t.Helper()
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}
			got := alerter.Alerts[i]
			poker.AssertScheduledAlert(t, got, want)
		})
	}
}
