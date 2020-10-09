package wallet

import "testing"


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


}