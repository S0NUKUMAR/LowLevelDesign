package models

type Player struct {
	Name string
	Role string

	Stats PlayerStats
}

type PlayerStats struct {
	RunsScored  int
	BallsFaced  int
	Fours       int
	Sixes       int
	Wickets     int
	OversBowled float64
	RunsGiven   int
	Catches     int
}

func NewPlayer(name, role string) *Player {
	return &Player{
		Name: name,
		Role: role,
	}
}
