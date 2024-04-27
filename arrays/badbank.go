package arrays

type Transactions struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transactions, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
