package status

import (
	"github.com/sonu/snakesladders/player"
)

type Mode string

const (
	NEW       Mode = "new"
	COMPLETED Mode = "completed"
	RUNNING   Mode = "running"
)

type Status struct {
	CurrentState Mode
	Winner       *player.Player
}

func (s Status) String() string {
	winnerName := "none"
	if s.Winner != nil {
		winnerName = s.Winner.Name
	}
	return "State: " + string(s.CurrentState) + ", Winner: " + winnerName
}
