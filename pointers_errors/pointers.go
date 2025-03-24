package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b *Bitcoin) val () int {
	return b.val()
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientBalance error = errors.New("insufficient balance to withdraw")


type Wallet struct {
	balance Bitcoin
}

type SuperWallet Wallet

func (w *Wallet) Deposit(val Bitcoin) {
	w.balance += val
	// pointers are implicitly dereferenced, so the below works too
	// w.balance += val 
}

func(w *Wallet) Withdraw(val Bitcoin) error {

	if w.balance < val {
		return ErrInsufficientBalance
	}

	w.balance -= val

	return nil
}

func (w *Wallet) Balance() Bitcoin{
	return w.balance
}