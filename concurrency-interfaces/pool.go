package main

import "strings"

type WorkerPool struct {
	poolNumber int
	in         (chan string)
	out        (chan string)
}

func (w *WorkerPool) Start() {
	for range w.poolNumber {
		go func() {
			for v := range w.in {
				w.out <- strings.ToUpper(v)
			}
		}()
	}
}

func (w *WorkerPool) Put(s string) {
	w.in <- s
}

func (w *WorkerPool) Stop() {
	close(w.in)
	close(w.out)
}

func (w *WorkerPool) Collect() chan string {
	return w.out
}
