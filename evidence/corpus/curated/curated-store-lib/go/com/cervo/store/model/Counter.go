package model

type Counter struct {
	value int
}

func NewCounter(value int) *Counter {
	return &Counter{value: value}
}
func (c *Counter) Add(delta int) {
	c.value = c.value + delta
}
func (c *Counter) Current() int {
	return c.value
}
