package pointers

import (
	"fmt"
	"github.com/pkg/errors"
)


type Bitcoin int

type Stringer interface {
	String() string
}

// Method function on the type Bitcoin, this satifies adherence to the Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// De-referencing of struct pointers is done by Go automically
//
//	return (*w).balance is the same as w.balance
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance when calling Deposit is %p \n", &w.balance)
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrWalletDoesntExist = errors.New("wallet does not exist")

// Wrapping Errors Messages
// given a custom error message - var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
// we can wrap this with further context - return errors.Wrap(ErrInsufficientFunds, "attempted to transfer more than balance")
// this will return a new error with the message:
//  	"attempted to transfer more than balance: cannot withdraw, insufficient funds"
// Unwrap the error message when testing using Cause() from the errors package
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w == nil {
		return errors.Wrap(ErrWalletDoesntExist, "attempted to withdraw from a wallet that doesn't exist")
	}
	if amount > w.balance {
		return errors.Wrap(ErrInsufficientFunds, "attempted to withdraw more than balance")
	}
	
	w.balance -= amount
	return nil
}

func (w *Wallet) Transfer(amount Bitcoin, destinationWallet *Wallet) error {
	if w == nil {
		return errors.Wrap(ErrWalletDoesntExist, "source wallet does not exist")
	}
	if destinationWallet == nil {
		return errors.Wrap(ErrWalletDoesntExist, "destination wallet does not exist")
	}
	
	if amount > w.balance {
		return errors.Wrap(ErrInsufficientFunds, "attempted to transfer more than balance")
	}
	
	w.balance -= amount
	destinationWallet.balance += amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
