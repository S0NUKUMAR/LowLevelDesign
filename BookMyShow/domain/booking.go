package domain

import (
	"time"
)

type BookingStatus string

var (
	BookingPending BookingStatus = "PENDING"
	BookingConfirmed BookingStatus = "CONFIRMED"
	BookingCancelled BookingStatus = "CANCELLED"
	BookingFailed BookingStatus = "FAILED"
)

type Booking struct {
	ID string
	ShowID string
	UserID string
	SeatIDs []string
	Amount float64
	Status BookingStatus
	CreatedAt time.Time
}

func NewBooking(id, showID, userID string, seatIDs []string, amount float64, createdAt time.Time) *Booking {
	return &Booking {
		ID:id,
		ShowID:showID,
		UserID:userID,
		SeatIDs:seatIDs,
		Amount:amount,
		CreatedAt:createdAt,
		Status:BookingPending,
	}
}
