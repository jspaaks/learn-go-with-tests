package wallet

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
