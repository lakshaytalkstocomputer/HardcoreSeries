package wallet

import (
	"testing"
)


func TestWallet(t *testing.T){
	// Test for adding amount in wallet
	// Test for checking Balance in wallet
	wallet := Wallet{}

	t.Run("deposit amount in wallet", func(t *testing.T) {
		got := wallet.Deposit(10)

		var expected error = nil

		if got!= expected{
			t.Errorf("got %v expected %v",got, expected)
		}
	})

	t.Run("Check balance", func(t *testing.T) {
		got := wallet.Balance()
		expected := 10.0

		if got != expected{
			t.Errorf("got %f expected %f",got, expected)
		}

	})


	t.Run("Withdraw Money", func(t *testing.T){
		err := wallet.Withdraw(5)
		if err != nil {
			t.Errorf("Found error while withdrawing")
		}

		got := wallet.Balance()
		expected := 5.0

		if got != expected{
			t.Errorf("got %f expected %f", got, expected)
		}
	})

	t.Run("Withdraw Money more than it is in Balance", func(t *testing.T) {
		err := wallet.Withdraw(100.0)

		if err == nil{
			t.Errorf("Error should not be null")
		}

	})


}