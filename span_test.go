package clog

import (
	"testing"
	"time"
)

func TestSpan_StartTransaction(t *testing.T) {
	s := &Span{}
	want := Transaction{name: "test"}
	if got := s.StartTransaction("test"); got.name != want.name {
		t.Errorf("StartTransaction() = %v, want %v", got, want)
	}
}

func TestSpan_FinishTransaction(t *testing.T) {
	s := &Span{}
	transaction := Transaction{name: "test", timestamp: time.Now()}
	s.FinishTransaction(transaction)

	if len(s.transactions) <= 0 {
		t.Errorf("FinishTransaction() must be adding a transaction to span")
	}

	if len(s.transactions) > 0 && s.transactions[0].duration <= 0 {
		t.Errorf("FinishTransaction() must be adding a transaction with duration, got %v", s.transactions[0])
	}
}
