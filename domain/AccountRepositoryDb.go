package domain

import (
	"banking/errs"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

const MsgAccountCreated = "Account created"

func (a AccountRepositoryDb) Create(customerId int64, t AccountType) *errs.AppError {

	insertQuery := "INSERT INTO account(customer_id, type, balance, created, status) VALUES(?, ?, 0, CURRENT_TIMESTAMP, 1)"

	_, err := a.db.Exec(insertQuery, customerId, t)

	if err != nil {
		return &errs.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &errs.AppError{Code: http.StatusOK, Message: MsgAccountCreated}
}

func NewAccountRepositoryDb() AccountRepositoryDb {
	return AccountRepositoryDb{db: NewSqlxDb()}
}

