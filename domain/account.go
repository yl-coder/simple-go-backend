package domain

import (
	"banking/errs"
	"errors"
)

type AccountType string

const (
	AccountSaving = "SAVING"
	AccountChecking = "CHECKING"
)

const INVALID_ENUM_TYPE = "Invalid enum type"

func (t AccountType) Validate() error {
	if t != AccountSaving && t != AccountChecking{
		return errors.New(INVALID_ENUM_TYPE)
	}
	return nil
}


type Account struct {
	id int64
	customerId int64
	t string
	balance float64
	status uint8
}

type AccountRepository interface {
	Create(customerId int64, t AccountType) *errs.AppError
}




