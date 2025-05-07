package counter

import "sync"

type Counter interface {
	Increment()
	GetValue() int
}

type counter struct {
	mutex sync.Mutex
	value int
}

func (c *counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *counter) GetValue() int {
	return c.value
}

func NewCounter() Counter {
	return &counter{}
}
