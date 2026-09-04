package domain

import "errors"


var (
	ErrSeatNotFound = errors.New("Seat not Found")
	ErrSeatNotAvailable = errors.New("Seat is not Available")
	ErrSeatNotHeldByUser = errors.New("Seat is not held by this User")
	ErrHoldExpired = errors.New("Seat Hold is Expired")
	ErrShowNotFound = errors.New("Seat Not Found")
	ErrBookingNotFound = errors.New("booking Not Found")
	ErrInvalidSeatSelection = errors.New("no Seat Selected")
	ErrPaymentFailed = errors.New("payment processing Failed")
)

	
