package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertCorrectBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertCorrectError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error but didn't get one.")
		}
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("Unexpectedly got an error.")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0}
		err := wallet.Deposit(10)
		assertNoError(t, err)
		assertCorrectBalance(t, wallet, Bitcoin(10))
	})

	t.Run("deposit a negative amount", func(t *testing.T) {
		const startingBalance = Bitcoin(0)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Deposit(-10)
		assertCorrectError(t, err, NegativeDepositError)
		assertCorrectBalance(t, wallet, startingBalance)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 100}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertCorrectBalance(t, wallet, Bitcoin(90))
	})

	t.Run("withdraw a negative amount", func(t *testing.T) {
		const startingBalance = Bitcoin(0)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(-10)
		assertCorrectError(t, err, NegativeWithdrawalError)
		assertCorrectBalance(t, wallet, startingBalance)
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(100)
		assertCorrectError(t, err, InsufficientFundsError)
		assertCorrectBalance(t, wallet, Bitcoin(20))
	})
}
