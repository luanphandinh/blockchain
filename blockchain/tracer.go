package blockchain

type Tracer interface {
	Trace(msg string)
	Tracef(msg string, args ...interface{})

	// Special function for trace carriage data
	// the msg return will have the format of `\r`
	TraceCarriagef(msg string, args ...interface{})
}

var tracer Tracer = &doNothingTracer{}

// Allow external caller to get some information inside blockchain package
// such as: Blocks being initialized, computed hash...
// all message from blockchain package are not ended with new line character ('\n')
func SetTracer(t Tracer) {
	tracer = t
}

type doNothingTracer struct{}

func (t *doNothingTracer) Trace(msg string)                               {}
func (t *doNothingTracer) Tracef(msg string, args ...interface{})         {}
func (t *doNothingTracer) TraceCarriagef(msg string, args ...interface{}) {}
