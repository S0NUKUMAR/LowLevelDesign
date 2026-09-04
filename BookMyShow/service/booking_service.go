package service

import (
	"time"
	"example.com/bookmyshow/domain"
	"example.com/bookmyshow/repository"
)

type BookingService struct {
	shows repository.ShowRepository
	bookings repository.BookingRepository
	payment PaymentGateway
}


func NewBookingService(shows repository.ShowRepository, bookings repository.BookingRepository, payment PaymentGateway) *BookingService {
	return &BookingService{
		shows:shows,
		bookings: bookings,
		payment: payment,
		}
}


func (b *BookingService) HoldSeats(showID, userID string, seatIDs []string) error {
	show, err := b.shows.FindById(showID)
	if err!=nil {
		return err 
	}

	return show.HoldSeat(userID, seatIDs, time.Now())
} 


func (b *BookingService) ReleaseHold(showID, userID string, seatIDs []string) error {
	show , err:= b.shows.FindById(showID)
	if err!= nil {
		return err
	}

	show.ReleaseSeats(userID,seatIDs)
	return nil
}



func (b *BookingService) ConfirmBooking(showID, userID string, seatIDs []string) (*domain.Booking,error){
	show, err:= b.shows.FindById(showID)
	if err!=nil {
		return nil, err
	}

	now:= time.Now()

	for _, id:= range seatIDs {
		state, err:= show.SeatState(id,now)
		if err!=nil{
			return nil, err
		}

		if state.Status!=domain.SeatHeld || state.HeldBy != userID {
			return nil, domain.ErrSeatNotHeldByUser
		}
	}

	amount := show.PriceFor(seatIDs)

	if _, err:= b.payment.Charge(userID, amount); err!=nil {
		return nil, err
	}

	bookingID:= nextID("booking")
	if err:= show.ConfirmSeats(userID, seatIDs, bookingID, time.Now()) ; err!=  nil {
		return nil, err	
	}

	booking:=domain.NewBooking(bookingID, showID, userID, seatIDs, amount, time.Now())
	booking.Status = domain.BookingConfirmed
	b.bookings.Save(booking)

	return booking, nil
}
