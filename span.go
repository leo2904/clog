package clog

import (
	"fmt"
	"sync"
	"time"
)

const (
	// IncrementalSpan if only add this tag when create the Span, then only count the transactions but not
	// calculate any duration
	IncrementalSpan = iota + 1

	// DurationSpan if only add this tag when create the Span, then only calculate the total duration of
	// the all transactions but not count them.
	DurationSpan
)

// Span represents an active, un-finished span in the Canonical Log Line system.
type Span struct {
	sync.RWMutex

	key          string
	flag         uint8
	transactions []Transaction
}

// Transaction represent a sequence of information exchange and related work.
type Transaction struct {
	name      string
	timestamp time.Time
	duration  time.Duration
}

// StartTransaction initialize a new transaction, this save the timestamp when transaction start.
func (s *Span) StartTransaction(name string) Transaction {
	return Transaction{
		name:      name,
		timestamp: time.Now(),
	}
}

// FinishTransaction calculate the duration of the all transaction period and add to the transactions span.
func (s *Span) FinishTransaction(t Transaction) {
	duration := time.Since(t.timestamp)
	t.duration = duration

	s.Lock()
	s.transactions = append(s.transactions, t)
	s.Unlock()
}

// String return a string with a representation of the transaction inside the span
// according to your type.
func (s *Span) String() string {
	s.Lock()
	var msg string

	if s.flag&DurationSpan != 0 {
		var d time.Duration
		for _, t := range s.transactions {
			d += t.duration
		}
		msg += fmt.Sprintf("%s_duration=%f ", s.key, d.Seconds())
	}
	if s.flag&IncrementalSpan != 0 {
		msg += fmt.Sprintf("%s_total=%d ", s.key, len(s.transactions))
	}
	s.Unlock()
	return msg
}
