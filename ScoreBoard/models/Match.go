package models

import (
	"time"

	"github.com/sksingh/enums"
)

type Match struct {
	TeamA *Team
	TeamB *Team

	MatchId string
	Format  enums.MatchFormat
	Status  enums.MatchStatus

	StartTime time.Time
	Innings   []Innings
}

func NewMatch(teamA *Team, teamB *Team) Match {
	return Match{
		TeamA: teamA,
		TeamB: teamB,
	}
}

func (m *Match) StartMatch() {
	m.Status = enums.Ongoing
	m.StartTime = time.Now()
}

func (m *Match) AddInning(inning Innings) {
	m.Innings = append(m.Innings, inning)
}
