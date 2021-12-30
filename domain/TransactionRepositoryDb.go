package domain

import (
	"banking/errs"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type TransactionRepositoryDb struct {
	db *sqlx.DB
}

func (a TransactionRepositoryDb) Create(accountId int64, amount float64, t TransactionType) *errs.AppError {
	tx, txErr := a.db.Begin()
	appError := createHelper(accountId, amount, t, txErr, tx)
	tx.Commit()
	return appError
}

func createHelper(accountId int64, amount float64, t TransactionType, txErr error, tx *sql.Tx) (*errs.AppError) {
	if txErr != nil {
		return errs.NewUnexpectedError("Unable to open commit")
	}

	selectQuery := "SELECT balance FROM account WHERE id = ? FOR UPDATE"

	row := tx.QueryRow(selectQuery, accountId)

	if row == nil {
		return errs.NewBadRequestError("Unable to find account")
	}

	var balance float64

	row.Scan(&balance)

	if t == TransactionDebit && balance-amount < 0 {
		return errs.NewBadRequestError("Insufficient amount in bank account")
	}

	insertQuery := "INSERT INTO transaction(account_id, type, amount, created) VALUES(?,?,?,CURRENT_TIMESTAMP)"

	_, insertErr := tx.Exec(insertQuery, accountId, t, amount)

	if insertErr != nil {
		return errs.NewUnexpectedError(insertErr.Error())
	}


	var newAmount float64

	if t == TransactionCredit {
		newAmount = balance + amount
	} else {
		newAmount = balance - amount
	}

	updateQuery := "UPDATE account set balance = ? WHERE id = ?"

	_, err := tx.Exec(updateQuery, newAmount, accountId)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return &errs.AppError{Code: http.StatusOK, Message: "Account created"}
}

func NewTransactionRepositoryDb() TransactionRepositoryDb {
	return TransactionRepositoryDb{db: NewSqlxDb()}
}

