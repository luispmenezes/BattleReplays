package main

import (
	battlereplays "github.com/luispmenezes/battle-replays/pkg"
	"log"
	"os"
)

func main() {
	f, err := os.Open("210511-182739.clientreplay")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	header, err := battlereplays.NewParser(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(header)
}
