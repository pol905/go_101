package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(wallet Wallet, shouldHave Bitcoin, t testing.TB) {
		t.Helper()
		balance := wallet.Balance()

		if balance != shouldHave {
			t.Errorf("Balance: %s, should have had: %s", balance, shouldHave)
		}
	}

	assertNoError := func(got error, t testing.TB) {
		t.Helper()

		if got != nil {
			t.Fatal("Wanted no Errors")
		}
	}

	assertError := func(got, want error, t testing.TB) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error")
		}

		if got.Error() != want.Error() {
			t.Errorf("Got: %q, Want: %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(wallet, Bitcoin(10), t)

	})

	t.Run("Withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(err, t)
		assertBalance(wallet, Bitcoin(0), t)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(230))

		assertBalance(wallet, startingBalance, t)

		assertError(err, ErrInsufficientFunds, t)

	})
}
