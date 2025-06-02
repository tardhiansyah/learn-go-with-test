package badbank

import (
	"testing"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		bob   = Account{Name: "Bob", Balance: 150}
		alice = Account{Name: "Alice", Balance: 200}
	)

	transactions := []Transaction{
		NewTransaction(riya, bob, 50),
		NewTransaction(bob, alice, 30),
		NewTransaction(alice, riya, 20),
	}

	newBalanceFor := func(account Account) float64 {
		newAccountBalance := NewBalanceFor(transactions, account).Balance
		return newAccountBalance
	}

	AssertEqual(t, newBalanceFor(riya), 70)
	AssertEqual(t, newBalanceFor(bob), 170)
	AssertEqual(t, newBalanceFor(alice), 210)
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
