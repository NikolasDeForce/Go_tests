package arrays

type Transactions struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transactions {
	return Transactions{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transactions) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transactions Transactions) Account {
	if transactions.From == a.Name {
		a.Balance -= transactions.Sum
	}
	if transactions.To == a.Name {
		a.Balance += transactions.Sum
	}
	return a
}
