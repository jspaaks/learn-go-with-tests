package bank

import "github.com/jspaaks/learn-go-with-tests/arrays-revisited/pkg/arrays"

func applyTransaction(transaction Transaction, a Account) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func ApplyTransactions(transactions []Transaction, a Account) Account {
	return arrays.Reduce(transactions, applyTransaction, a)
}
