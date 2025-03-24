package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertWalletBalanceIsCorrect := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("wallet: %#v ,got %s want %s", wallet, got, want)
		}
	}

	assertErrorIsNotNil := func(t testing.TB, err error){
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}
	}

	assertErrorIsNil := func(t testing.TB, err error){
		if err != nil {
			t.Errorf("wanted error to be nil, but got error")
		}
	}


	assertErrorMessageIsExpected := func(t testing.TB, got string, want string){
		if got != want {
			t.Errorf("got %s, want %s", got ,want)
		}
	}

	t.Run("Test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
	
		assertWalletBalanceIsCorrect(t, wallet, Bitcoin(10))
	})
	
	t.Run("Test Valid Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertErrorIsNil(t, err)
		assertWalletBalanceIsCorrect(t, wallet, Bitcoin(5))
			
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
	
		assertWalletBalanceIsCorrect(t, wallet, startingBalance)
		assertErrorIsNotNil(t, err)
		assertErrorMessageIsExpected(t, err.Error(),ErrInsufficientBalance.Error())
	})

	

}