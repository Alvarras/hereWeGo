package dasar

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(c *customer, t transaction) error {
	if t.transactionType == transactionDeposit {
		c.balance += t.amount
		return nil
	} else if t.transactionType == transactionWithdrawal {
		if c.balance < t.amount {
			return errors.New("insufficient funds")
		}
		c.balance -= t.amount
		return nil
	}
	return errors.New("unknown transaction type")
}
