package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.v[key] += 2
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func main() {
	s := SafeCounter{v: make(map[string]int)}

	for i := 2; i < 1000; i++ {
		go s.Inc("numbers")
	}
	time.Sleep(time.Second)
	fmt.Println(s.Value("numbers"))
}
