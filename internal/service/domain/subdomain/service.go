package subdomain

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
)

type Service struct {
	entities      map[uint64]*Subdomain
	entitiesIndex []uint64
	indexSync     sync.RWMutex
	seriesId      uint64
}

func (s *Service) Describe(ID uint64) (*Subdomain, error) {
	s.indexSync.RLock()
	defer s.indexSync.RUnlock()
	if m, ok := s.entities[ID]; ok {
		return m, nil
	}
	return nil, fmt.Errorf("not found model with id=%d", ID)
}

func (s *Service) List(cursor uint64, limit uint64) ([]Subdomain, error) {
	if limit == 0 {
		return []Subdomain{}, nil
	}

	pos := 0
	if cursor > 0 {
		pos = s.findIndexPosition(cursor) + 1
	}
	if pos >= len(s.entitiesIndex) {
		return []Subdomain{}, fmt.Errorf("no more elements")
	}

	slice := make([]Subdomain, 0, limit)
	foundElements := uint64(0)
	for i := pos; i < len(s.entitiesIndex); i++ {
		ID := s.entitiesIndex[i]
		slice = append(slice, *s.entities[ID])
		foundElements++

		if foundElements >= limit {
			break
		}
	}

	return slice, nil
}

func (s *Service) Create(subdomain Subdomain) (uint64, error) {
	s.indexSync.Lock()
	defer s.indexSync.Unlock()

	id := s.getNextId()
	subdomainNew := recreateSubdomain(id, subdomain)

	s.entities[id] = subdomainNew
	s.entitiesIndex = append(s.entitiesIndex, id)

	return id, nil
}

func (s *Service) Update(ID uint64, subdomain Subdomain) error {
	updatingSubdomain, err := s.Describe(ID)
	if err != nil {
		return err
	}

	s.indexSync.Lock()
	defer s.indexSync.Unlock()

	updatingSubdomain.fillFrom(subdomain)

	return nil
}

func (s *Service) Remove(ID uint64) (bool, error) {
	_, err := s.Describe(ID)
	if err != nil {
		return false, err
	}
	s.indexSync.Lock()
	defer s.indexSync.Unlock()

	delete(s.entities, ID)
	pos := s.findIndexPosition(ID)
	if pos == len(s.entitiesIndex) {
		s.entitiesIndex = s.entitiesIndex[:pos]
	} else if pos == 0 {
		s.entitiesIndex = s.entitiesIndex[1:]
	} else {
		s.entitiesIndex = append(s.entitiesIndex[:pos], s.entitiesIndex[pos+1:]...)
	}

	return true, nil
}

func (s *Service) findIndexPosition(ID uint64) int {
	return sort.Search(len(s.entitiesIndex), func(i int) bool {
		return s.entitiesIndex[i] >= ID
	})
}

func NewService() ServiceInterface {
	service := &Service{
		entities:      make(map[uint64]*Subdomain, 0),
		entitiesIndex: make([]uint64, 0),
		seriesId:      0,
	}
	return service
}

func (s *Service) getNextId() uint64 {
	return atomic.AddUint64(&s.seriesId, 1)
}
