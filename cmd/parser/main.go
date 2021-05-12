package main

import (
	BattleReplays "BattleReplays/pkg"
	"log"
	"os"
)

func main() {
    f, err := os.Open("210511-182739.clientreplay")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = BattleReplays.NewParser(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("##############################################################")

	f2, err := os.Open("210511-183422.clientreplay")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	_, err = BattleReplays.NewParser(f2)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("##############################################################")

	f3, err := os.Open("210509-211435.clientreplay")
	if err != nil {
		log.Fatal(err)
	}
	defer f3.Close()

	_, err = BattleReplays.NewParser(f3)
	if err != nil {
		log.Fatal(err)
	}
}

