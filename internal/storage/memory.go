package storage

import (
	"errors"
	"math/rand"
	"sync"
	"test/internal/models"
)

type MemoryStore struct {
	sync.RWMutex
	quotes []models.Quote
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		quotes: []models.Quote{},
		nextID: 1,
	}
}

func (s *MemoryStore) Add(q models.Quote) models.Quote {
	s.Lock()
	defer s.Unlock()

	q.ID = s.nextID
	s.nextID++
	s.quotes = append(s.quotes, q)

	return q
}

func (s *MemoryStore) GetAll() []models.Quote {
	s.RLock()
	defer s.RUnlock()

	return s.quotes
}

func (s *MemoryStore) GetByAuthor(author string) []models.Quote {
	s.RLock()
	defer s.RUnlock()

	var result []models.Quote
	for _, q := range s.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}

	return result
}

func (s *MemoryStore) GetRandom() (models.Quote, error) {
	s.RLock()
	defer s.RUnlock()

	if len(s.quotes) == 0 {
		return models.Quote{}, errors.New("нет доступных цитат")
	}

	return s.quotes[rand.Intn(len(s.quotes))], nil
}

func (s *MemoryStore) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	for i, q := range s.quotes {
		if q.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return nil
		}
	}

	return errors.New("цитата не найдена")
}
