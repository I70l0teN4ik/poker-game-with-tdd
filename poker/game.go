package poker

import (
	"time"
)

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewGame(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{alerter, store}
}

func (g *TexasHoldem) Start(numOfPlayers int) {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	interval := time.Duration(numOfPlayers + 5) * time.Minute
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += interval
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}