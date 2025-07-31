package main

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c Counter) GetValue() int {
	return c.value
}
