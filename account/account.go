package account

import (
	"errors"
	"fmt"
	"time"
)

//Account is an interface that wraps the common behavior for accounts.
type Account interface {
	Deposit(amount Money) error
	Withdraw(amount Money) error
}

//Transfer transfers money from one account to another. Returns the root cause in case of an error.
func Transfer(fromAcct Account, toAcct Account, amount Money) error {
	if err := fromAcct.Withdraw(amount); err == nil {
		if depErr := toAcct.Deposit(amount); depErr == nil {
			fmt.Printf("Transfered %f from %s to %+v", amount, fromAcct, toAcct)
		} else {
			//return the root cause
			return depErr
		}
	} else {
		//return the root cause
		return err
	}
	return nil
}

//Money type.
type Money float64

//SavingsAccount encapsulates the state of a savings account.
type SavingsAccount struct {
	InterestRate float32
	MinBalance   Money

	Num      string
	Name     string
	OpenDate time.Time
	Balance  Money
}

//CheckingAccount encapsulates the state of a checking account.
type CheckingAccount struct {
	TransactionFee   float32
	OverDraftEnabled bool

	Num      string
	Name     string
	OpenDate time.Time
	Balance  Money
}

//OpenSavingsAccount is the Initializing constructor for SavingsAccount
func OpenSavingsAccount(no string, name string, openingDate time.Time) *SavingsAccount {
	a := SavingsAccount{
		Num:          no,
		Name:         name,
		OpenDate:     openingDate,
		InterestRate: 0.9,
		MinBalance:   15.0,
	}
	return &a
}

//OpenCheckingAccount is the initializing constructor for CheckingAccount
func OpenCheckingAccount(no string, name string, openingDate time.Time, overdraftFlag bool) *CheckingAccount {
	return &CheckingAccount{
		Num:              no,
		Name:             name,
		OpenDate:         openingDate,
		TransactionFee:   0.15,
		OverDraftEnabled: overdraftFlag,
	}
}

//Deposit adds specified amount to existing balance of savings account. Returns nil on successful deposit.
func (acct *SavingsAccount) Deposit(amount Money) error {
	fmt.Printf("Depositing %f \n", amount)
	acct.Balance = acct.Balance + amount
	return nil
}

//Withdraw removes specified amount from existing balance of savings account. Returns nil on successful withdrawal.
func (acct *SavingsAccount) Withdraw(amount Money) error {
	//check for min balance invariant
	if acct.Balance-amount < acct.MinBalance {
		return errors.New("not enough money to withdraw")
	}
	acct.Balance = acct.Balance - amount
	return nil
}

//Deposit adds specified amount to existing balance of checking account. Returns nil on successful deposit.
func (acct *CheckingAccount) Deposit(amount Money) error {
	fmt.Printf("Depositing %f \n", amount)
	acct.Balance = acct.Balance + amount
	return nil
}

//Withdraw removes specified amount from existing balance of checking account. Returns nil on successful withdrawal.
func (acct *CheckingAccount) Withdraw(amount Money) error {
	//check the overdraft invariant
	if acct.Balance < amount && !acct.OverDraftEnabled {
		return errors.New("Not enough money to withdraw")
	}
	acct.Balance = acct.Balance - amount
	return nil
}
