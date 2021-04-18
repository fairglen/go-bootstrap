package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://dev.to/quii/learn-go-by-writing-tests-pointers-and-errors-2kp6

func TestWallet(t *testing.T){
	t.Run("deposit", func(t *testing.T){
		// given
		anAmount := Bitcoin(10)
		wallet := Wallet{}
		//	when
		wallet.Deposit(anAmount)
		balance := wallet.Balance()
		// then
		require.Equal(t, anAmount, balance, "deposit: got %s wanted %s", balance, anAmount)
	})
	t.Run("withdraw", func(t *testing.T){
		// given
		anAmount := Bitcoin(10)
		wallet := Wallet{balance: 2 * anAmount}
		//	when
		wallet.Withdraw(anAmount)
		balance := wallet.Balance()
		// then
		require.Equal(t, anAmount, balance, "withdraw: got %s wanted %s", balance, anAmount)
	})
	t.Run("insufficient funds", func(t *testing.T){
		// given
		anAmount := Bitcoin(10)
		wallet := Wallet{}
		//	when
		err := wallet.Withdraw(anAmount)
		// then
		require.Error(t, InsufficientFundsError, err)
	})
}
