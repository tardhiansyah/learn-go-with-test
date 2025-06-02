package badbank

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  sum,
	}
}

func NewBalanceFor(transactions []Transaction, account Account) Account {
	return Reduce(transactions, applyTransaction, account)
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - float64(t.Sum)
		}
		if t.To == name {
			return currentBalance + float64(t.Sum)
		}
		return currentBalance
	}

	return Reduce(transactions, adjustBalance, 0.0)
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue

	for _, item := range collection {
		result = f(result, item)
	}

	return result
}

func applyTransaction(account Account, t Transaction) Account {
	if t.From == account.Name {
		account.Balance -= t.Sum
	} else if t.To == account.Name {
		account.Balance += t.Sum
	}
	return account
}
