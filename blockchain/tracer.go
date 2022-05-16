package blockchain

import "context"

type Tracer interface {
	Trace(ctx context.Context, msg string)
	Tracef(ctx context.Context, msg string, args ...interface{})

	// Special function for trace carriage data
	// the msg return will have the format of `\r`
	TraceCarriagef(ctx context.Context, msg string, args ...interface{})
}

var tracer Tracer = &doNothingTracer{}

// Allow external caller to get some information inside blockchain package
// such as: Blocks being initialized, computed hash...
// all messages from blockchain package are not ended with new line character ('\n')
func SetTracer(t Tracer) {
	tracer = t
}

type doNothingTracer struct{}

func (t *doNothingTracer) Trace(ctx context.Context, msg string)                               {}
func (t *doNothingTracer) Tracef(ctx context.Context, msg string, args ...interface{})         {}
func (t *doNothingTracer) TraceCarriagef(ctx context.Context, msg string, args ...interface{}) {}
