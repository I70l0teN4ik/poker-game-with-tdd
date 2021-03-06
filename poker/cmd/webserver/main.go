package main

import (
	"log"
	"net/http"
	"tdd/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)
	server, err := poker.NewPlayerServer(store, game)
	if nil != err {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
