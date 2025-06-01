package badbank

import (
	"testing"
)

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Alice",
			To:   "Bob",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Bob",
			Sum:  50,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Bob"), 150)
	AssertEqual(t, BalanceFor(transactions, "Alice"), -100)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -50)
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
