package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	elves, err := processElves("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("max elf", top(elves))
	fmt.Println("top 3 elves", topN(elves, 3))
}

type elf struct {
	inventory []string
}

func topN(elves []elf, n int) int {
	var elfCals []int
	for _, e := range elves {
		elfCals = append(elfCals, e.calories())
	}
	// Sort elfCals.
	slices.SortFunc(elfCals, func(x, y int) bool {
		return y < x
	})

	topN := elfCals[0:n]
	total := 0
	for _, i := range topN {
		total += i
	}
	return total
}

func top(elves []elf) int {
	maxElf := elves[0]
	for _, elf := range elves[1:] {
		if elf.calories() > maxElf.calories() {
			maxElf = elf
		}
	}
	return maxElf.calories()
}

func (e *elf) calories() int {
	// Pattern 3: Finding the sum of elements in an array.
	sum := 0
	for _, c := range e.inventory {
		v, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		sum += v
	}
	return sum
}

// processElves process the elves calories from an input file and returns
// a slice of elves.
func processElves(fp string) ([]elf, error) {
	fpath, err := filepath.Abs(fp)
	if err != nil {
		return nil, err
	}
	// Pattern 1: Processing a file's data.
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	var elves []elf

	// Process the data.
	parsed := strings.Split(string(data), "\n")
	lastIdx := 0
	for i, s := range parsed {
		if s == "" {
			elf := elf{inventory: parsed[lastIdx : i+1]}
			elves = append(elves, elf)
			lastIdx = i
		}
	}

	return elves, nil
}
