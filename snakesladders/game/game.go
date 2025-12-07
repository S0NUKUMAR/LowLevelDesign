package game

import (
	"fmt"

	"github.com/sonu/snakesladders/board"
	"github.com/sonu/snakesladders/dice"
	"github.com/sonu/snakesladders/player"
	"github.com/sonu/snakesladders/status"
)

type Game struct {
	Status  *status.Status
	Player  []*player.Player
	Snakes  []board.Snake
	Ladders []board.Ladder
	Dice    dice.RandomDice
}

func NewGame(players []*player.Player, snakes []board.Snake, ladders []board.Ladder, dice dice.RandomDice) *Game {
	return &Game{
		Status:  &status.Status{CurrentState: status.NEW},
		Player:  players,
		Snakes:  snakes,
		Ladders: ladders,
		Dice:    dice,
	}
}

func (g *Game) PlayTurn(p *player.Player) {
	num := g.Dice.Roll()
	fmt.Printf("%s rolled %d\n", p.Name, num)

	if p.Position+num <= 100 {
		old := p.Position
		p.Position += num
		p.Position = board.ApplyBoardRules(p.Position, g.Snakes, g.Ladders)
		fmt.Printf("Player %s moved from %d ---> %d\n", p.Name, old, p.Position)

	}

	if p.Position == 100 {
		g.Status.CurrentState = status.COMPLETED
		g.Status.Winner = p
	}

}
