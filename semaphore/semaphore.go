package semaphore

import (
	"sync"
)

type Semaphore struct {
	l     sync.Mutex
	count int
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		count: size,
	}
}

// Return true if semaphore decremented.
func (s *Semaphore) Wait() bool {
	s.l.Lock()
	defer s.l.Unlock()
	if s.count <= 0 {
		return false
	}

	s.count--
	return true
}

func (s *Semaphore) Signal() {
	s.l.Lock()
	defer s.l.Unlock()
	s.count++
}
