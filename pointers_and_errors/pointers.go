package pointers

import "fmt"

type Wallet struct {
    // Remember, In Go variables,types,functins etc
    // that start with lowercase are private OUTSIDE the package it's defined in
    balance int
}

func (w *Wallet) Deposit(money int) {
   fmt.Printf("memory address of balance in Deposit method: '%v' \n", &w.balance)
   w.balance += money
}

// w is the Reciever variable here
// In go, when you call a function/mehod the arguments are copied.
// The 'w' is copy of whatever the caller invoked
// To keep state, we can fix this with pointers.
// Pointers let us point to a value and let us update it
// Lets take a pointer to the wallet so we can udpate it.
// *Wallet is read as "a pointer to a wallet value"

// Technically we don't need the pointer here since we're not updating state,
// we could use a copy of the wallet instead.
func (w *Wallet) Balance() int {
    // Go doesn't require us to deference the pointer but we could with:
    // (*w).balance
    // Struct pointers are automatically deferenced, they're referred to as struct pointers
    return w.balance
}
