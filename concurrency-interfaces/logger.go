package main

import (
	"bytes"
	"fmt"
)

type Logger interface {
	Log(message string)
	Error(err error)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func (c ConsoleLogger) Error(err error) {
	fmt.Println(err)
}

type FileLogger struct {
	buf *bytes.Buffer
}

func (f *FileLogger) Log(message string) {
	b := []byte(message)
	f.buf.Write(b)
}

func (f *FileLogger) Error(err error) {
	f.buf.Write(fmt.Appendln([]byte{}, err))
}
