package wallet


type Wallet struct{
	balance float64
}

func (w *Wallet) Deposit(amount float64) error{
	w.balance += amount
	return nil
}

func (w Wallet) Balance()float64{
	return w.balance
}