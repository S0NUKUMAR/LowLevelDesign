package domain

type Screen struct{
	ID string
	Name string
	Seats []*Seat
}

type SeatType string

var (
	SILVER SeatType = "SILVER"
	GOLD SeatType = "GOLD"
)

type Seat struct {
	ID string
	Row string
	Number int
	Type SeatType
}


func NewScreen(id, name string, seats []*Seat) *Screen {
	return &Screen {
		ID: id, 
		Name :name,
		Seats: seats,
	}
}


