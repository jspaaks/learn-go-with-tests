package wallet

import "errors"

var InsufficientFundsError = errors.New("Insufficient funds available")
var NegativeWithdrawalError = errors.New("Can't withdraw a negative number")
var NegativeDepositError = errors.New("Can't deposit a negative number")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return NegativeDepositError
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	if amount <= 0 {
		return NegativeWithdrawalError
	}
	w.balance -= amount
	return nil
}
