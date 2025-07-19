package enums

type MatchFormat string

const (
	T20  MatchFormat = "T20"
	ODI  MatchFormat = "ODI"
	Test MatchFormat = "TEST"
)

type MatchStatus string

const (
	NotStarted MatchStatus = "NOT_STARTED"
	Ongoing    MatchStatus = "ONGOING"
	Paused     MatchStatus = "PAUSED"
	Finished   MatchStatus = "FINISHED"
)

type BallType string

const (
	Normal BallType = "NORMAL"
	NoBall BallType = "NO_BALL"
	Bye    BallType = "BYE"
	LegBye BallType = "LEG_BYE"
	Wide   BallType = "WIDE"
)

type Dismissaltype string

const (
	None    Dismissaltype = "NONE"
	Bowled  Dismissaltype = "BOWLED"
	RunOut  Dismissaltype = "RUN_OUT"
	Caught  Dismissaltype = "CAUGHT"
	Stumped Dismissaltype = "STUMPED"
)
