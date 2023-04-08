package main

import (
	"fmt"
	"log"

	"day01/internal/elf"
)

func main() {
	elves, err := elf.ProcessElves("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	top := elf.Top(elves)
	fmt.Println("top", top)

	top3 := elf.TopN(elves, 3)
	fmt.Println("top3", top3)
}
