package main

import (
	BattleReplays "BattleReplays/pkg"
	"log"
	"os"
)

func main() {
	f, err := os.Open("210510-170214.clientreplay")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = BattleReplays.NewParser(f)
	if err != nil {
		log.Fatal(err)
	}
}
