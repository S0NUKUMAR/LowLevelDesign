package models

type Team struct {
	Name         string
	Players      []Player
	Captain      Player
	ViceCaptain  Player
	WicketKeeper Player
}
