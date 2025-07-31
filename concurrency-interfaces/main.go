package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			for range 1000 {
				c.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Value: %d\n", c.GetValue())
}

func polimorficLogger() {
	c := ConsoleLogger{}
	var b bytes.Buffer
	f := FileLogger{
		buf: &b,
	}
	ProcessData(c, "to console")
	ProcessData(&f, "to buffer")
	fmt.Printf("From buffer: %s\n", b.String())
	b.Reset()
	tc := TimestampLogger{c}
	tf := TimestampLogger{&f}
	ProcessData(tc, "to console")
	ProcessData(tf, "to buffer")
	fmt.Printf("From buffer: %s\n", b.String())
}

func ProcessData(logger Logger, data string) {
	logger.Log(data)
}
