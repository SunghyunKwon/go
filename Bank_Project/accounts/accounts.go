package accounts

import "errors"

// Account structure
type Account struct {
	owner   string
	balance int
}

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't widthdraw you are poor")
	}

	a.balance -= amount
	return nil
}

// ChangeOwner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner
func (a Account) Owner() string {
	return a.owner
}
