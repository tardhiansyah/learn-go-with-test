package badbank

type Transaction struct {
	From string
	To   string
	Sum  int
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= float64(t.Sum)
		}
		if t.To == name {
			balance += float64(t.Sum)
		}
	}
	return balance
}
