package repository


import (
	"sync"
	"example.com/bookmyshow/domain"
)

type BookingRepository interface {
	Save (booking *domain.Booking)
	FindByID(id string) (*domain.Booking, error)
	FindByUser(userID string) []*domain.Booking
}

type InMemoryBookingRepository struct {
	mu sync.RWMutex
	bookings map[string]*domain.Booking
}

func NewInMemoryBookingRepository() *InMemoryBookingRepository{
	return &InMemoryBookingRepository{
		bookings:make(map[string]*domain.Booking),
	} 
}

func (r *InMemoryBookingRepository) Save(booking *domain.Booking) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.bookings[booking.ID] = booking
}

func (r *InMemoryBookingRepository)FindByID(id string) (*domain.Booking, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	booking, ok:= r.bookings[id]; 
	if !ok {
		return nil, domain.ErrBookingNotFound
	}

	return booking , nil
}


func (r *InMemoryBookingRepository) FindByUser(userID string) []*domain.Booking {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.Booking

	for _, b:= range r.bookings {
		if b.UserID == userID{
			result = append(result,b)
		}
	}

	return result
}
