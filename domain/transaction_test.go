package domain

import (
	"testing"
)

func Test_Given_TransactionTypeInvalid_When_MethodValidate_ThenReturnError(t *testing.T) {

	var transactionType TransactionType
	transactionType = "Invalid"

	err := transactionType.Validate()

	if err == nil {
		t.Error("Err should be not nil, as transaction type is invalid")
	}
}

func Test_Given_TransactionTypeDebit_When_MethodValidate_ThenReturnNoError(t *testing.T) {

	var transactionType TransactionType
	transactionType = TransactionDebit

	err := transactionType.Validate()

	if err != nil {
		t.Error("Err should be nil, as transaction type is valid")
	}
}


func Test_Given_TransactionTypeCredit_When_MethodValidate_ThenReturnNoError(t *testing.T) {

	var transactionType TransactionType
	transactionType = TransactionCredit

	err := transactionType.Validate()

	if err != nil {
		t.Error("Err should be nil, as transaction type is valid")
	}
}