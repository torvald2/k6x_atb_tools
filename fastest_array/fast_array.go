package fastest_array

import (
	"sync"
)

var stringArrayObject stringArray
var once_d sync.Once

const (
	DATE_FORMAT string = "2006-01-02"
)

type StringArrayCreator struct{}

type stringArray struct {
	data []string
}

func (r *StringArrayCreator) Create(data []string) *stringArray {
	once_d.Do(func() {
		stringArrayObject = stringArray{data}
	})
	return &stringArrayObject
}

func (i *stringArray) GetData(index int) string {

	return i.data[index]
}
