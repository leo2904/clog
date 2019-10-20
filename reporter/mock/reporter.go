package mock

import (
	"github.com/friendsofgo/clog/reporter"
)

type Reporter struct {
	SendInvoked  bool
	CloseInvoked bool

	Sendfn  func(fmtLine string, severity reporter.Severity)
	Closefn func() error
}

func (r *Reporter) Send(fmtLine string, severity reporter.Severity) {
	r.SendInvoked = true
	r.Sendfn(fmtLine, severity)

}
func (r *Reporter) Close() error {
	r.CloseInvoked = true
	return r.Closefn()
}
