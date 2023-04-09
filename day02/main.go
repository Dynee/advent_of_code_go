package main

import (
	"day02/internal/rps"
	"fmt"
	"log"
)

func main() {
	playerScore, err := rps.StartGame("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("player score", playerScore)
}
