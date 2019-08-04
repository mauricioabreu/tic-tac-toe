package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board 	[9]Mark
	moves 	[]int
}

type Mark struct {
	symbol string
}

func (m Mark) value() int {
	switch m.symbol {
	case "X":
		return 1
	case "O":
		return 2
	default:
		return 0
	}
}

var combinations = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{6, 4, 2},
}

func NewGame() *Game{
	return &Game{}
}

func (g *Game) Draw(symbol string, spot int) error {
	if g.board[spot].symbol != "" {
		return errors.New("spot already marked")
	}
	g.board[spot] = Mark{symbol}
	g.moves = append(g.moves, spot)
	return nil
}

func (g *Game) Winner() string {
	for _, combination := range combinations {
		result := sum(g.board[combination[0]].value(), g.board[combination[1]].value(), g.board[combination[2]].value()) 
		switch result {
		case 3:
			return "X"
		case 6:
			return "O"
		default:
			continue
		}
	}
	return ""
}

func sum(values ...int) int{
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	fmt.Println("Starting game...")
	game := NewGame()
	game.Draw("X", 0)
	game.Winner()
}