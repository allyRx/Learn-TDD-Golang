package main

import (
	"errors"
	"fmt"
)


type Bitcoin int

type Wallet struct{
	balance		Bitcoin
}


func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin){
	fmt.Printf("adress of balance in deposit is %p \n", &w.balance)
	w.balance +=  amount
}

func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	
	if amount > w.balance{
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil
}
func main(){
	w := Wallet{}

	 w.Deposit(Bitcoin(10))
	fmt.Println(w.Balance())
}