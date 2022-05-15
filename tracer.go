package main

import "fmt"

type simpleTracer struct{}

func (t *simpleTracer) Trace(msg string) {
	fmt.Println(msg)
}

func (t *simpleTracer) Tracef(msg string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(msg, args...))
}

func (t *simpleTracer) TraceCarriagef(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}
