package board

import "fmt"

type Snake struct {
	Head int
	Tail int
}

type Ladder struct {
	Start int
	End   int
}

func ApplyBoardRules(pos int, snakes []Snake, ladders []Ladder) int {
	for _, s := range snakes {
		if pos == s.Head {
			fmt.Printf("Oops! Hit a snake 🐍 down to %d\n", s.Tail)
			return s.Tail
		}
	}

	for _, l := range ladders {
		if pos == l.Start {
			fmt.Printf("Yay! Climbed a ladder 🪜  up to %d\n", l.End)
			return l.End
		}
	}
	return pos
}
