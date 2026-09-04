package repository

import (
	"example.com/bookmyshow/domain"
	"sync"
)


type ShowRepository interface {
	Save(show *domain.Show)
	FindById(id string) (*domain.Show, error)
	FindAll()[]*domain.Show
}


type InMemoryShowRepository struct {
	mu sync.RWMutex
	shows map[string]*domain.Show
}


func NewInMemoryShowRepository() *InMemoryShowRepository {
	return &InMemoryShowRepository {
		shows: make(map[string]*domain.Show),
	}
}

func (r *InMemoryShowRepository) Save(show *domain.Show){
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shows[show.ID] = show
}


func(r *InMemoryShowRepository) FindById(id string) (*domain.Show, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	show, ok := r.shows[id]
	if !ok {
		return nil, domain.ErrShowNotFound
	}

	return show, nil
}

func (r *InMemoryShowRepository) FindAll() []*domain.Show {
	r.mu.RLock()
	defer r.mu.RUnlock()

	all:= make([]*domain.Show,0,len(r.shows))
	for _, s:= range r.shows{
		all = append(all,s)
	}
	return all
}



