package pointers

import (
	"fmt"
	"testing"
	"github.com/pkg/errors"
)

// Assertion helper functions
// Balance, no error, wrapped error
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertWrappedError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
    if errors.Cause(got) != want {
        t.Errorf("got %v, want %v", got, want)
    }
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

// Wallet test
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(90))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertWrappedError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

	t.Run("Withdraw from nil wallet", func(t *testing.T) {
		var wallet *Wallet // create a nil pointer
		err := wallet.Withdraw(Bitcoin(20))
		assertWrappedError(t, err, ErrWalletDoesntExist)
	})
}
func TestWalletExtensionTests(t *testing.T) {
	tests := []struct {
		name               string
		sourceWallet       *Wallet
		destinationWallet  *Wallet
		amount             Bitcoin
		expectedSourceBalance  Bitcoin
		expectedDestBalance    Bitcoin
		expectedError        error
	}{
		{
			name:              "Transfer sufficient funds",
			sourceWallet:      &Wallet{balance: Bitcoin(100)},
			destinationWallet: &Wallet{balance: Bitcoin(50)},
			amount:            Bitcoin(40),
			expectedSourceBalance: Bitcoin(60),
			expectedDestBalance:   Bitcoin(90),
			expectedError:       nil,
		},
		{
			name:              "Transfer insufficient funds",
			sourceWallet:      &Wallet{balance: Bitcoin(20)},
			destinationWallet: &Wallet{balance: Bitcoin(50)},
			amount:            Bitcoin(100),
			expectedSourceBalance: Bitcoin(20),
			expectedDestBalance:   Bitcoin(50),
			expectedError:       ErrInsufficientFunds,
		},
		{
			name:              "Transfer from a nil wallet",
			sourceWallet:      nil,
			destinationWallet: &Wallet{balance: Bitcoin(50)},
			amount:            Bitcoin(10),
			expectedSourceBalance: Bitcoin(0),
			expectedDestBalance:   Bitcoin(50),
			expectedError:       ErrWalletDoesntExist,
		},
		{
			name:              "Transfer to a nil wallet",
			sourceWallet:      &Wallet{balance: Bitcoin(100)},
			destinationWallet: nil,
			amount:            Bitcoin(10),
			expectedSourceBalance: Bitcoin(100),
			expectedDestBalance:   Bitcoin(0),
			expectedError:       ErrWalletDoesntExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sourceWallet.Transfer(tt.amount, tt.destinationWallet)

			if tt.expectedError != nil {
				assertWrappedError(t, err, tt.expectedError)
			} else {
				assertNoError(t, err)
			}

			if tt.sourceWallet != nil {
				assertBalance(t, *tt.sourceWallet, tt.expectedSourceBalance)
			}

			if tt.destinationWallet != nil {
				assertBalance(t, *tt.destinationWallet, tt.expectedDestBalance)
			}
		})
	}
}
