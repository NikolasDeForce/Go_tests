package arrays

type Transactions struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transactions, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transactions) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}

		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}
