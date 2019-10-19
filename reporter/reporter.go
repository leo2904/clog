package reporter

// Reporter interface can be used to provide the Canonical Log Line with custom implementations
// to publish log lines.
type Reporter interface {
	// Send canonical log line to the reporter.
	Send(fmtLine string)
	// Close close the reporter.
	Close() error
}

type noopReporter struct{}

// NewNoop returns a no-op Reporter implementation.
func NewNoop() Reporter {
	return &noopReporter{}
}

func (r noopReporter) Send(fmtLine string) {}
func (r noopReporter) Close() error        { return nil }