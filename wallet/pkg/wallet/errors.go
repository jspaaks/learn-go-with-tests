package wallet

import "errors"

// Error issued when a user tries to withdraw more money from the wallet than they have available.
var InsufficientFundsError = errors.New("Insufficient funds available")

// Error issued when the user tries to withdraw a negative amount.
var NegativeWithdrawalError = errors.New("Can't withdraw a negative number")

// Error issued when the user tries to deposit a negative amount.
var NegativeDepositError = errors.New("Can't deposit a negative number")
