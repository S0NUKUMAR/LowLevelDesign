package scoreboard

import (
	"fmt"

	"github.com/sksingh/models"
)

func ShowScore(match *models.Match) {
	fmt.Printf("Match: %s vs %s | Format: %s | Status: %s\n",
		match.TeamA.Name, match.TeamB.Name, match.Format, match.Status)

	if len(match.Innings) == 0 {
		fmt.Println("No innings yet.")
		return
	}

	for i, inning := range match.Innings {
		fmt.Printf("Inning %d: %s - %d/%d in %.1f overs\n",
			i+1, inning.BattingTeam.Name, inning.TotalRuns, inning.TotalWickets, inning.OversPlayed)
	}
}

func ShowDetailedScoreboard(inning models.Innings) {
	fmt.Printf("\n📋 %s Innings\n", inning.BattingTeam.Name)
	fmt.Printf("Total: %d/%d in %.1f overs\n", inning.TotalRuns, inning.TotalWickets, inning.OversPlayed)
	fmt.Println("\nBatting\t\t\tR  B  4s 6s  SR")
	for _, p := range inning.BattingTeam.Players {
		if p.Stats.BallsFaced == 0 {
			continue
		}
		sr := float64(p.Stats.RunsScored) / float64(p.Stats.BallsFaced) * 100
		fmt.Printf("%-16s %3d %3d %3d %3d %6.2f\n", p.Name, p.Stats.RunsScored, p.Stats.BallsFaced, p.Stats.Fours, p.Stats.Sixes, sr)
	}

	fmt.Println("\nBowling\t\t\tO   R  W  Econ")
	for _, p := range inning.BowlingTeam.Players {
		if p.Stats.OversBowled == 0 {
			continue
		}
		econ := float64(p.Stats.RunsGiven) / p.Stats.OversBowled
		fmt.Printf("%-16s %3.1f %3d %2d %5.2f\n", p.Name, p.Stats.OversBowled, p.Stats.RunsGiven, p.Stats.Wickets, econ)
	}
}
