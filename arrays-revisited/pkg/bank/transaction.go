package bank

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from Account, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}
