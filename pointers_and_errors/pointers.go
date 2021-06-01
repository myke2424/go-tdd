package pointers

import (
	"errors"
	"fmt"
)

// We can declare methods on our bitcoin type. This is useful for adding domain specific functionality.
// Lets implement 'Stringer' on bitcoin.
// This interface lets you define how  your type is printed when using with %s format strings
// This is analogous to pythons __repr__ or __str__
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// Remember, In Go variables,types,functins etc
	// that start with lowercase are private OUTSIDE the package it's defined in
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("memory address of balance in Deposit method: '%v' \n", &w.balance)
	w.balance += amount
}

// var keyword allows use to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
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
func (w *Wallet) Balance() Bitcoin {
	// Go doesn't require us to deference the pointer but we could with:
	// (*w).balance
	// Struct pointers are automatically deferenced, they're referred to as struct pointers
	return w.balance
}

// Generally you should keep your method reciever types to be the same for consistancy

// Go lets you create new types from existing ones... Lets create a bitcoin int type.


// WRAP UP
// Go copies values when you pass them to functions/methods so if you need mutate state
// you'll need a pointer to that "thing" you want to change

// The fact go takes a copy of the value is nice but sometimes you want to pass by reference (pointer) for larger data
// Or something you would have one instance of (like a database connection)

// Pointers can be nil
// When a function returns a pointer to something, you need to check if it's nil so you dont raise a runtime exception
// Nil is useful when you want to describe a value that could be missing

// Errors are the way to signify failue when calling a function/method

// Creating new types from exisiting ones are useful for addng more domain specific meaning to values
