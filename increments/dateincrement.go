package increments

import (
	"sync"
	"time"
)

var dateIncrementObject dateIncrement
var once sync.Once

const (
	DATE_FORMAT string = "2006-01-02"
)

type DateIncrementCreator struct{}

type dateIncrement struct {
	CurrentDate time.Time
	Increment   int
	mu          sync.Mutex
}

func (r *DateIncrementCreator) New(start string, step int) *dateIncrement {
	once.Do(func() {
		t, err := time.Parse(DATE_FORMAT, start)
		if err != nil {
			panic("Date format must be 2006-01-02 ")
		}
		dateIncrementObject = dateIncrement{CurrentDate: t, Increment: step}
	})
	return &dateIncrementObject
}

func (i *dateIncrement) Get() time.Time {
	i.mu.Lock()
	defer func() {
		i.CurrentDate.AddDate(0, 0, i.Increment)
		i.mu.Unlock()
	}()
	return i.CurrentDate
}
