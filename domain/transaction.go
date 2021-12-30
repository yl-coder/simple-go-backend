package domain

import (
	"banking/errs"
	"errors"
	"time"
)

type TransactionType string

const (
	TransactionDebit = "DEBIT"
	TransactionCredit = "CREDIT"
)

func (t TransactionType) Validate() error {
	if t != TransactionDebit && t != TransactionCredit {
		return errors.New(INVALID_ENUM_TYPE)
	}
	return nil
}

type Transaction struct {
	id int64
	accountId int64
	t string
	amount float64
	created time.Time
}

type TransactionRepository interface {
	Create(accountId int64, amount float64, t TransactionType) *errs.AppError
}




