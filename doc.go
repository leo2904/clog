/*
Package clog is a structured logger for Go which follows the guidelines collected by Stripe in his article: https://stripe.com/es/blog/canonical-log-lines.
You can start to use the clog package follow the next example:
  package main
  import (
    "github.com/friendsofgo/clog"
  )
  func main() {
    logger := clog.New()
	line, _ := logger.NewLineOnContext(context.TODO())
	line.AddTags(clog.String("auth_type", "api"), clog.Int("rate_quota", 100))
	line.Send()
  }
Output:
  [2019-10-27T23:23:45+01:00] canonical-log-line auth_type=api rate_quota=100
For a full guide visit https://github.com/friendsofgo/clog
*/
package clog
