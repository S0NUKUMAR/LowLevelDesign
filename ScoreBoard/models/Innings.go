package models

type Innings struct {
	BattingTeam  *Team
	BowlingTeam  *Team
	TotalRuns    int
	TotalWickets int
	OversPlayed  float64
	Overs        []Over
}

func (inn *Innings) AddBall(ball Ball) {
	lastOverIndex := len(inn.Overs) - 1
	if lastOverIndex < 0 || len(inn.Overs[lastOverIndex].Balls) == 6 {
		inn.Overs = append(inn.Overs, Over{
			OverNumber: len(inn.Overs) + 1,
			Balls:      []Ball{},
		})
		lastOverIndex++
	}
	inn.Overs[lastOverIndex].Balls = append(inn.Overs[lastOverIndex].Balls, ball)
	inn.TotalRuns += ball.Runs
	if ball.Dismissal != "NONE" {
		inn.TotalWickets++
	}
	inn.OversPlayed = float64(len(inn.Overs)-1) + float64(len(inn.Overs[lastOverIndex].Balls))/6
}
