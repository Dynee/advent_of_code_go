package rps_test

import (
	"day02/internal/rps"
	"testing"
)

func TestStartGame(t *testing.T) {
	wantScore := 15

	gotScore, err := rps.StartGame("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if wantScore != gotScore {
		t.Fatal("scores do not match.")
	}
}
