package storages

import (
	"sync"
	"time"
)

type Storage[T any] struct {
	Mutex    *sync.Mutex
	storage  map[string]Entry[T]
	interval time.Duration
	done     chan struct{}
}

type Entry[T any] struct {
	Data       T
	Expiration time.Time
}

func New[T any](interval time.Duration) (s *Storage[T]) {
	s = &Storage[T]{
		Mutex:    &sync.Mutex{},
		storage:  make(map[string]Entry[T]),
		interval: interval,
		done:     make(chan struct{}),
	}

	go s.gc()

	return
}

func (s *Storage[T]) Get(key string) (value T) {
	if len(key) <= 0 {
		return
	}

	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	v := s.storage[key]

	if time.Now().After(v.Expiration) {
		delete(s.storage, key)
		return
	}

	value = v.Data
	return
}

func (s *Storage[T]) Set(key string, value T, expiration time.Duration) {
	if len(key) <= 0 || expiration < time.Second {
		return
	}

	s.Mutex.Lock()

	s.storage[key] = Entry[T]{
		Data:       value,
		Expiration: time.Now().Add(expiration),
	}

	s.Mutex.Unlock()
}

func (s *Storage[T]) Delete(key string) (err error) {
	if len(key) <= 0 {
		return
	}

	s.Mutex.Lock()

	delete(s.storage, key)

	s.Mutex.Unlock()
	return
}

func (s *Storage[T]) Reset() (err error) {
	s.Mutex.Lock()

	s.storage = make(map[string]Entry[T])

	s.Mutex.Unlock()
	return
}

func (s *Storage[T]) Close() (err error) {
	s.done <- struct{}{}
	return
}

func (s *Storage[T]) gc() {
	t := time.NewTicker(s.interval)
	defer t.Stop()

	for {
		select {
		case <-s.done:
			return
		case t := <-t.C:
			s.Mutex.Lock()

			for id, v := range s.storage {
				if t.After(v.Expiration) {
					delete(s.storage, id)
				}
			}

			s.Mutex.Unlock()
		}
	}
}
