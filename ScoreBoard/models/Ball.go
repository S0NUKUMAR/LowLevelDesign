package models

import "github.com/sksingh/enums"

type Ball struct {
	BallNumber int
	BatsMan    *Player
	Bowler     *Player
	Type       enums.BallType      //NO_BALL , WIDE , NORMAL
	Dismissal  enums.Dismissaltype //NONE , CAUGHT , RUN_OUT , STUMPED
	Runs       int
}
