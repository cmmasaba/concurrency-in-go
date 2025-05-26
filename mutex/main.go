package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	// Only one goroutine at a time can increment a key in the map
	s.v[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Only one goroutine at a time can access and return the value of a key from the map
	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}