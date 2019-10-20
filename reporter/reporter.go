package reporter

// Severity use severity in your lines for helping to decide which messages are send in
// each reporter
type Severity uint8

const (
	// SeverityInfo mark lines with this level for lines about information.
	SeverityInfo = iota + 1
	// SeverityError mark lines with this level for lines about controlled errors.
	SeverityError
	// SeverityCritical mark lines with this level for lines about critical errors.
	SeverityCritical
)

// Reporter interface can be used to provide the Canonical Log Line with custom implementations
// to publish log lines.
type Reporter interface {
	// Send canonical log line to the reporter.
	Send(fmtLine string, severity Severity)
	// Close close the reporter.
	Close() error
}

type noopReporter struct{}

// NewNoop returns a no-op Reporter implementation.
func NewNoop() Reporter {
	return &noopReporter{}
}

func (r noopReporter) Send(fmtLine string, severity Severity) {}
func (r noopReporter) Close() error                           { return nil }
