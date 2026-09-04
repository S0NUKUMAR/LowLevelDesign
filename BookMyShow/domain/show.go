package domain


import (
	"time"
	"sync"
)

type SeatStatus string

const(
	SeatAvailable SeatStatus = "AVAILABLE"
	SeatHeld SeatStatus = "HELD"
	SeatBooked SeatStatus = "BOOKED"
)


const DefaultHoldDuration = 5 * time.Second


type ShowSeat struct {
	SeatID string
	Status SeatStatus
	HeldBy string
	HoldExpiry time.Time
	BookingID string
}


type Show struct {
	ID string
	Movie *Movie	
	Screen *Screen
	Theatre *Theatre
	StartTime time.Time
	EndTime time.Time
	Price map[SeatType]float64


	mu sync.RWMutex
	seats map[string]*ShowSeat
}

func NewShow(id string, movie *Movie, theatre *Theatre, screen *Screen, start, end time.Time, price map[SeatType]float64) *Show {
	seats:= make(map[string]*ShowSeat,len(screen.Seats))
	for _, s:= range screen.Seats {
		seats[s.ID] = &ShowSeat{
			SeatID: s.ID,
			Status: SeatAvailable,
		}
	}
	return &Show{
		ID :id,
		Movie: movie,
		Theatre: theatre,
		StartTime :start,
		Screen:screen,
		EndTime: end,
		Price: price,
		seats: seats,
	}
}

func (s *Show)releaseExpiredLocked(seat *ShowSeat, now time.Time) {
	if seat.Status == SeatHeld && now.After(seat.HoldExpiry) {
		seat.Status = SeatAvailable
		seat.HeldBy = ""
		seat.HoldExpiry = time.Time{}
	}
}


func (s *Show) HoldSeat(userID string, seatIDs []string, now time.Time) error {
	if len(seatIDs) == 0 {
		return ErrInvalidSeatSelection
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for _ , id:= range seatIDs {
		seat, ok := s.seats[id]
		if !ok {
			return ErrSeatNotFound
		}

		s.releaseExpiredLocked(seat,now)
		if seat.Status != SeatAvailable {
			return ErrSeatNotAvailable
		}
	}

	expiry:= now.Add(DefaultHoldDuration)
	for _, id := range seatIDs {
		seat := s.seats[id]
		seat.Status = SeatHeld
		seat.HeldBy = userID
		seat.HoldExpiry = expiry
	}

	return nil
}

func (s *Show) ConfirmSeats(userID string, seatIDs []string, bookingID string, now time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, id:= range seatIDs{
		seat, ok := s.seats[id]
		if !ok {
			return ErrSeatNotFound
		}
		s.releaseExpiredLocked(seat,now)
		if seat.Status != SeatHeld || seat.HeldBy != userID {
			return ErrSeatNotHeldByUser
		}
	}

	for _ , id := range seatIDs {
		seat := s.seats[id]
		seat.Status = SeatBooked
		seat.HeldBy = ""
		seat.HoldExpiry = time.Time{}
		seat.BookingID = bookingID	
	}

	return nil
}


func (s *Show) ReleaseSeats(userID string, seatIDs []string){
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, id := range seatIDs {
		seat,ok := s.seats[id]
		if !ok || seat.Status != SeatHeld || seat.HeldBy != userID{
			continue
		}

		seat.Status = SeatAvailable
		seat.HeldBy = ""
		seat.HoldExpiry = time.Time{}
	}
}

func (s *Show) SeatTypeOf(seatID string) (SeatType, bool) {
	for _ , seat := range s.Screen.Seats{
		if seat.ID == seatID {
			return seat.Type, true
		}	
	}
	return "", false
}

func (s *Show) PriceFor(seatIDs []string) float64 {
	var total float64
	for _, id:= range seatIDs{
		if t, ok:= s.SeatTypeOf(id); ok {
			total+=s.Price[t]
		}
	}

	return total
}

func (s *Show) SeatState(seatID string, now time.Time) (ShowSeat, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	seat, ok := s.seats[seatID]
	if !ok {
		return ShowSeat{}, ErrSeatNotFound
	}
	s.releaseExpiredLocked(seat, now)
	return *seat, nil
}

func (s *Show) AvailableSeats(now time.Time) []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var ids []string
	for id, seat := range s.seats {
		s.releaseExpiredLocked(seat, now)
		if seat.Status == SeatAvailable {
			ids = append(ids, id)
		}
	}
	return ids
}
