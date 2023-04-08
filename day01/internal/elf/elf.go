package elf

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

// Each elf has an inventory where they store calories.
type Elf struct {
	Inventory []string
}

// Calories calculates the total amount of calories for an elf
func (e *Elf) Calories() int {
	sum := 0
	for _, c := range e.Inventory {
		v, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		sum += v
	}
	return sum
}

// Top returns the elf with the most calories.
func Top(elves []Elf) int {
	maxElf := elves[0]
	for _, elf := range elves[1:] {
		if elf.Calories() > maxElf.Calories() {
			maxElf = elf
		}
	}
	return maxElf.Calories()
}

// TopN returns the top N elves with the most calories from elves.
func TopN(elves []Elf, n int) int {
	var elfCals []int
	for _, e := range elves {
		elfCals = append(elfCals, e.Calories())
	}

	// Sort elfcals.
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

// ProcessElves processes an input file and returns a slice of elves or an error.
func ProcessElves(fp string) ([]Elf, error) {
	fpath, err := filepath.Abs(fp)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	var elves []Elf

	// Process the data.
	parsed := strings.Split(string(data), "\n")
	lastIdx := 0
	for i, s := range parsed {
		if s == "" {
			elf := Elf{Inventory: parsed[lastIdx : i+1]}
			elves = append(elves, elf)
			lastIdx = i
		}
	}

	return elves, nil
}
