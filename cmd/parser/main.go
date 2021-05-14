package main

import (
	BattleReplays "BattleReplays/pkg"
	"log"
	"os"
)

func main() {
	f, err := os.Open("Battlerite-Example-Replay_1.4.clientreplay")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = BattleReplays.NewParser(f)
	if err != nil {
		log.Fatal(err)
	}
}
