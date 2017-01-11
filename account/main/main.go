package main

import (
	"fmt"
	acct "github.com/yaravind/learn-go/account"
	"reflect"
	"time"
)

func main() {
	aliceAcct := acct.OpenSavingsAccount("12345", "Alice", time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC))
	fmt.Println("Alice's account =", aliceAcct)
	aliceAcct.Deposit(acct.Money(100.0))
	fmt.Println("Alice's account (after deposit) =", aliceAcct)
	if err := aliceAcct.Withdraw(acct.Money(10)); err != nil {
		fmt.Println(err)
		fmt.Println(reflect.TypeOf(err))
	} else {
		fmt.Println("Alice's account (after withdrawl) =", aliceAcct)
	}
	const overDraft = false
	bobAcct := acct.OpenCheckingAccount("98765", "Bob", time.Date(1997, time.April, 03, 0, 0, 0, 0, time.UTC), overDraft)
	fmt.Println("\nBob's account =", bobAcct)
	bobAcct.Deposit(acct.Money(100.0))
	fmt.Println("Bob's account (after deposit) =", bobAcct)
	if err := bobAcct.Withdraw(acct.Money(77)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bob's account (after withdrawl) =", bobAcct)
	}

	fmt.Println("\nTransfering 10.0 from Alice to Bob's acct")
	if err := acct.Transfer(aliceAcct, bobAcct, acct.Money(76.0)); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Alice's account =", aliceAcct)
	fmt.Println("Bob's account =", bobAcct)
}

type AccountRepository interface {
	FindById(id int64)
	Delete(id int64)
}
type SavingsRepository interface {
	Store(a acct.SavingsAccount)
}
type CheckingRepository interface {
	Store(a acct.CheckingAccount)
}
