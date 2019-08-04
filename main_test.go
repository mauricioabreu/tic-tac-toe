package main

import (
	"testing"
)

func TestDrawOnEmptySpot(t *testing.T) {
	game := NewGame()
	err := game.Draw("X", 0)
	if err != nil {
		t.Error("marking an empty spot should not fail")
	}
}

func TestDrawOnAlreadyMarkedSpot(t *testing.T) {
	game := NewGame()
	game.Draw("X", 0)
	err := game.Draw("X", 0)
	if err == nil {
		t.Error("marking an already marked spot should fail")
	}
}

func TestVictoryWithThreeMarks(t *testing.T) {
	game := NewGame()
	game.Draw("X", 0)
	game.Draw("X", 1)
	game.Draw("X", 2)
	winner := game.Winner()
	if winner != "X" {
		t.Errorf("winner of this game should be X, but got %s", winner)
	}
}