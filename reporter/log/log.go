package log

import (
	"log"
	"os"

	"github.com/friendsofgo/clog/reporter"
)

type logReporter struct {
	logger *log.Logger
}

// NewReporter returns a new log reporter.
func NewReporter(l *log.Logger) reporter.Reporter {
	if l == nil {
		// if nil logger if received then we initialize the logger to standard type
		l = log.New(os.Stderr, "", 0)
	}
	return &logReporter{logger: l}
}

func (r logReporter) Send(fmtLine string, severity reporter.Severity) {
	r.logger.Print(fmtLine)
}

func (r logReporter) Close() error {return nil}


