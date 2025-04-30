package wallet

import "errors"

// Error issued when a user tries to withdraw more money from the wallet than they have available.
var InsufficientFundsError = errors.New("Insufficient funds available")

// Error issued when the user tries to withdraw a negative amount.
var NegativeWithdrawalError = errors.New("Can't withdraw a negative number")

// Error issued when the user tries to deposit a negative amount.
var NegativeDepositError = errors.New("Can't deposit a negative number")

// Defintion of a wallet.
type Wallet struct {
	balance Bitcoin
}

// Returns the current balance of the receiver wallet w.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit amount into the receiver wallet w.
func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return NegativeDepositError
	}
	w.balance += amount
	return nil
}

// Withdraw amount from the receiver wallet w.
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
