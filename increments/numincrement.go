package increments

import (
	"sync"
)

var numIncrementObject numIncrement
var once sync.Once

type NumIncrementCreator struct{}

type numIncrement struct {
	CurrentNumber int
	Increment     int
	mu            sync.Mutex
}

func (r *NumIncrementCreator) Create(start, step int) *numIncrement {
	once.Do(func() {
		numIncrementObject = numIncrement{CurrentNumber: start, Increment: step}
	})
	return &numIncrementObject
}

func (i *numIncrement) Get() int {
	i.mu.Lock()
	defer func() {
		i.CurrentNumber = i.CurrentNumber + i.Increment
		i.mu.Unlock()
	}()
	return i.CurrentNumber
}
