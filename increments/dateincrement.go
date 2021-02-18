package increments

import (
	"sync"
	"time"
)

var dateIncrementObject dateIncrement
var once_d sync.Once

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
	once_d.Do(func() {
		t, err := time.Parse(DATE_FORMAT, start)
		if err != nil {
			panic("Date format must be 2006-01-02 ")
		}
		dateIncrementObject = dateIncrement{CurrentDate: t, Increment: step}
	})
	return &dateIncrementObject
}

func (i *dateIncrement) Get() string {
	i.mu.Lock()
	defer func() {
		i.CurrentDate = i.CurrentDate.AddDate(0, 0, i.Increment)
		i.mu.Unlock()
	}()
	return i.CurrentDate.Format(DATE_FORMAT)
}
