package elf_test

import (
	"day01/internal/elf"
	"testing"
)

func TestTop(t *testing.T) {
	t.Parallel()

	wantTop := 400

	elves, err := elf.ProcessElves("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := elf.Top(elves)
	if wantTop != got {
		t.Fatal(err)
	}
}

func TestTopN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		n        int
		calories int
	}{
		{"top 3", 3, 900},
		{"top 2", 2, 700},
		{"top 1", 1, 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elves, err := elf.ProcessElves("./testdata/input.txt")
			if err != nil {
				t.Fatal(err)
			}

			got := elf.TopN(elves, tt.n)
			if tt.calories != got {
				t.Fatal(err)
			}
		})
	}
}
