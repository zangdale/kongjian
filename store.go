package kongjian

import (
	"sync"
)

type StoreInter[T any] interface {
	Add(v T) error
	Exist(v T) (bool, error)
	Delete(v T) error
}

var _ StoreInter[int64] = (*StoreInt64)(nil)

type StoreInt64 struct {
	sync.Mutex
	count map[int64]uint64
}

func NewStoreInt64() *StoreInt64 {
	return &StoreInt64{
		Mutex: sync.Mutex{},
		count: make(map[int64]uint64),
	}
}

func (s *StoreInt64) Add(v int64) error {
	s.Lock()
	defer s.Unlock()
	_, ok := s.count[v]
	if ok {
		s.count[v]++
	} else {
		s.count[v] = 1
	}
	return nil
}

func (s *StoreInt64) Exist(v int64) (bool, error) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.count[v]
	return ok, nil
}

func (s *StoreInt64) Delete(v int64) error {
	s.Lock()
	defer s.Unlock()
	_, ok := s.count[v]
	if ok {
		s.count[v]--
	}
	return nil
}
