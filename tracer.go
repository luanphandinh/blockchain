package main

import (
	"context"
	"fmt"
)

type simpleTracer struct{}

func (t *simpleTracer) Trace(ctx context.Context, msg string) {
	fmt.Println(msg)
}

func (t *simpleTracer) Tracef(ctx context.Context, msg string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(msg, args...))
}

func (t *simpleTracer) TraceCarriagef(ctx context.Context, msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}
