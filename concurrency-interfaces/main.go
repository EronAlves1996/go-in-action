package main

import (
	"bytes"
	"fmt"
)

func main() {
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
