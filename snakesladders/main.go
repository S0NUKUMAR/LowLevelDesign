package main

import (
	"fmt"
	"time"

	"github.com/sonu/snakesladders/board"
	"github.com/sonu/snakesladders/dice"
	"github.com/sonu/snakesladders/game"
	"github.com/sonu/snakesladders/player"
	"github.com/sonu/snakesladders/status"
)

func main() {
	p1 := &player.Player{Name: "Sonu"}
	p2 := &player.Player{Name: "Monu"}

	snakes := []board.Snake{{98, 11}, {50, 34}, {75, 5}, {34, 2}}
	ladders := []board.Ladder{{3, 22}, {5, 8}, {20, 29}, {27, 56}, {72, 91}}
	dice := dice.RandomDice{}

	g := game.NewGame([]*player.Player{p1, p2}, snakes, ladders, dice)
	g.Status.CurrentState = status.RUNNING

	turn := 0
	for g.Status.CurrentState != status.COMPLETED {
		g.PlayTurn(g.Player[turn])
		turn = 1 - turn
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Game Over!", g.Status)
}
