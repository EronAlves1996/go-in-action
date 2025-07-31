package main

import (
	"bytes"
	"fmt"
	"time"
)

type Logger interface {
	Log(message string)
	Error(err error)
}

type TimestampLogger struct {
	Logger
}

func (t *TimestampLogger) now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t TimestampLogger) Log(message string) {
	t.Logger.Log(fmt.Sprintf("[%s] %s", t.now(), message))
}

func (t TimestampLogger) Error(err error) {
	fmt.Printf("[%s] %s\n", t.now(), err)
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
