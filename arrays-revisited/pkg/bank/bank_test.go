package bank_test

import (
	"testing"
	"github.com/jspaaks/learn-go-with-tests/arrays-revisited/pkg/assertions"
	"github.com/jspaaks/learn-go-with-tests/arrays-revisited/pkg/bank"
)

func TestBank(t *testing.T) {
	riya := bank.Account{Name: "Riya", Balance: 100}
	chris := bank.Account{Name: "Chris", Balance: 75}
	adil := bank.Account{Name: "Adil", Balance: 200}
	transactions := []bank.Transaction{
		bank.NewTransaction(chris, riya, 100),
		bank.NewTransaction(adil, chris, 25),
	}
	newBalanceFor := func(account bank.Account) float64 {
		return bank.ApplyTransactions(transactions, account).Balance
	}
	assertions.AssertEqual(t, newBalanceFor(riya), 200)
	assertions.AssertEqual(t, newBalanceFor(chris), 0)
	assertions.AssertEqual(t, newBalanceFor(adil), 175)
}
