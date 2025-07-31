package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	w := WorkerPool{
		poolNumber: 5,
		in:         make(chan string, 100),
		out:        make(chan string, 100),
	}
	var wg sync.WaitGroup
	for range 70 {
		wg.Add(1)
		w.Put("aksfhaksfakAAakshajpqwiorwqwr")
	}
	w.Start()

	go func() {
		wg.Wait()
		w.Stop()
	}()

	for v := range w.Collect() {
		wg.Done()
		fmt.Println(v)
	}
}

func concurrentCounter() {
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
