package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sksingh/enums"
	"github.com/sksingh/models"
	"github.com/sksingh/scoreboard"
)

func createPlayers(names []string) []models.Player {
	var players []models.Player
	for _, name := range names {
		players = append(players, models.Player{Name: name})
	}
	return players
}

func nextBatsman(players []models.Player, used map[int]bool) *models.Player {
	for i := 0; i < len(players); i++ {
		if !used[i] {
			used[i] = true
			return &players[i]
		}
	}
	return nil
}

// Helper to simulate a single ball
func playBall(
	inning *models.Innings,
	striker **models.Player,
	nonStriker **models.Player,
	bowler *models.Player,
	usedBatsmen map[int]bool,
	ballNum *int,
	wickets *int,
	target int,
	isChasing bool,
) (endOfInnings bool) {
	outcome := rand.Intn(10)
	runs := 0
	dismissal := enums.None

	if outcome <= 6 {
		runs = outcome
	} else {
		dismissal = enums.Bowled
		(*wickets)++
	}

	ballObj := models.Ball{
		BallNumber: *ballNum,
		BatsMan:    *striker,
		Bowler:     bowler,
		Type:       enums.Normal,
		Runs:       runs,
		Dismissal:  dismissal,
	}
	(*ballNum)++
	inning.AddBall(ballObj)

	(*striker).Stats.RunsScored += runs
	(*striker).Stats.BallsFaced++
	bowler.Stats.RunsGiven += runs
	bowler.Stats.OversBowled += 1.0 / 6.0
	if dismissal != enums.None {
		bowler.Stats.Wickets++
		*striker = nextBatsman(inning.BattingTeam.Players, usedBatsmen)
		if *striker == nil {
			return true // All out
		}
		return false
	}

	if runs%2 == 1 {
		*striker, *nonStriker = *nonStriker, *striker
	}

	if isChasing && inning.TotalRuns >= target {
		return true // Target reached
	}

	return false
}

func simulateInning(battingTeam *models.Team, bowlingTeam *models.Team, maxOvers int) models.Innings {
	inning := models.Innings{
		BattingTeam: battingTeam,
		BowlingTeam: bowlingTeam,
	}

	usedBatsmen := make(map[int]bool)
	striker := nextBatsman(battingTeam.Players, usedBatsmen)
	nonStriker := nextBatsman(battingTeam.Players, usedBatsmen)
	wickets := 0
	bowlerIndices := []int{6, 7, 8, 9, 10}
	ballNum := 1
	for over := 1; over <= maxOvers && wickets < 10; over++ {
		bowler := &bowlingTeam.Players[bowlerIndices[(over-1)%len(bowlerIndices)]]
		for ball := 1; ball <= 6; ball++ {
			if wickets >= 10 {
				break
			}
			end := playBall(&inning, &striker, &nonStriker, bowler, usedBatsmen, &ballNum, &wickets, 0, false)
			if end {
				break
			}
		}
		striker, nonStriker = nonStriker, striker
	}
	return inning
}

func simulateSecondInning(battingTeam *models.Team, bowlingTeam *models.Team, target int, maxOvers int) models.Innings {
	inning := models.Innings{
		BattingTeam: battingTeam,
		BowlingTeam: bowlingTeam,
	}

	usedBatsmen := make(map[int]bool)
	striker := nextBatsman(battingTeam.Players, usedBatsmen)
	nonStriker := nextBatsman(battingTeam.Players, usedBatsmen)
	wickets := 0
	bowlerIndices := []int{6, 7, 8, 9, 10}
	ballNum := 1
	for over := 1; over <= maxOvers && wickets < 10 && inning.TotalRuns < target; over++ {
		bowler := &bowlingTeam.Players[bowlerIndices[(over-1)%len(bowlerIndices)]]
		for ball := 1; ball <= 6; ball++ {
			if wickets >= 10 || inning.TotalRuns >= target {
				break
			}
			end := playBall(&inning, &striker, &nonStriker, bowler, usedBatsmen, &ballNum, &wickets, target, true)
			if end {
				break
			}
		}
		striker, nonStriker = nonStriker, striker
	}
	return inning
}

func main() {
	rand.Seed(time.Now().UnixNano())

	indiaPlayers := createPlayers([]string{
		"Rohit Sharma", "Shubman Gill", "Virat Kohli", "Shreyas Iyer", "KL Rahul",
		"Hardik Pandya", "Ravindra Jadeja", "Shardul Thakur", "Kuldeep Yadav", "Mohammed Shami", "Jasprit Bumrah",
	})
	australiaPlayers := createPlayers([]string{
		"David Warner", "Travis Head", "Steve Smith", "Marnus Labuschagne", "Glenn Maxwell",
		"Marcus Stoinis", "Pat Cummins", "Mitchell Starc", "Josh Hazlewood", "Adam Zampa", "Alex Carey",
	})

	teamIndia := &models.Team{Name: "India", Players: indiaPlayers, Captain: indiaPlayers[2], WicketKeeper: indiaPlayers[4]}
	teamAus := &models.Team{Name: "Australia", Players: australiaPlayers, Captain: australiaPlayers[2], WicketKeeper: australiaPlayers[10]}

	match := &models.Match{
		MatchId:   "INDvsAUS2025",
		TeamA:     teamIndia,
		TeamB:     teamAus,
		Format:    enums.ODI,
		Status:    enums.NotStarted,
		StartTime: time.Now(),
	}

	match.StartMatch()

	// Inning 1 - India bats
	inning1 := simulateInning(teamIndia, teamAus, 50)
	match.AddInning(inning1)

	// Inning 2 - Australia chases
	inning2 := simulateSecondInning(teamAus, teamIndia, inning1.TotalRuns+1, 50)
	match.AddInning(inning2)

	scoreboard.ShowScore(match)
	scoreboard.ShowDetailedScoreboard(inning1)
	scoreboard.ShowDetailedScoreboard(inning2)
	printMatchResult(inning1, inning2)
}

func printMatchResult(inning1, inning2 models.Innings) {
	fmt.Println("\n🏁 Match Result:")
	fmt.Printf("%s: %d/%d in %.1f overs\n", inning1.BattingTeam.Name, inning1.TotalRuns, inning1.TotalWickets, inning1.OversPlayed)
	fmt.Printf("%s: %d/%d in %.1f overs\n", inning2.BattingTeam.Name, inning2.TotalRuns, inning2.TotalWickets, inning2.OversPlayed)

	if inning2.TotalRuns > inning1.TotalRuns {
		fmt.Printf("%s won by %d wickets\n", inning2.BattingTeam.Name, 10-inning2.TotalWickets)
	} else if inning2.TotalRuns < inning1.TotalRuns {
		fmt.Printf("%s won by %d runs\n", inning1.BattingTeam.Name, inning1.TotalRuns-inning2.TotalRuns)
	} else {
		fmt.Println("Match Tied")
	}
}
