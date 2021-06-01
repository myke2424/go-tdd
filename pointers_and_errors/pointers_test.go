package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(10)
		want := Bitcoin(10) // This is our bitcoin type mapped to int

		assertBalance(t, wallet, want)
		fmt.Printf("memory address of balance in test is: '%v' \n", &wallet.balance)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)

	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, 5)
		assertError(t, err, ErrInsufficientFunds.Error())
	})

}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	// t.Fatal will stop the test if it is called. We dont want to make any assertions on a error if there isnt one.
	if got == nil {
		t.Fatal("didn't get an error but wanted get one")
	}

	if got.Error() != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	balance := wallet.Balance()

	if balance != want {
		t.Errorf("got '%s' but want '%s'", balance, want)
	}
}
