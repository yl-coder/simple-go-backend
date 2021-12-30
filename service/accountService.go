package service

import (
	"banking/domain"
	"banking/errs"
)

//go:generate mockgen -source=accountService.go -destination=../mocks/service/mockAccountService.go -package=service
type AccountService interface {
	CreateAccount(customerId int64, t domain.AccountType) *errs.AppError
	CreateTransaction(accountId int64, amount float64, t domain.TransactionType) *errs.AppError
}

type AccountServiceImpl struct {
	accountRepo domain.AccountRepository
	transactionRepo domain.TransactionRepository
}

func (a AccountServiceImpl) CreateTransaction(accountId int64, amount float64, t domain.TransactionType) *errs.AppError {
	err := t.Validate()

	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	createErr := a.transactionRepo.Create(accountId, amount, t)

	return createErr
}

func (a AccountServiceImpl) CreateAccount(customerId int64, t domain.AccountType) *errs.AppError {
	validateErr := t.Validate()

	if validateErr != nil {
		return errs.NewBadRequestError(validateErr.Error())
	}

	err := a.accountRepo.Create(customerId, t)
	return err
}

func NewAccountServiceImpl(accountRepo domain.AccountRepository, transactionRepo domain.TransactionRepository) AccountService {
	return AccountServiceImpl{accountRepo: accountRepo, transactionRepo: transactionRepo}
}


