package domain

import (
	"banking/errs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func(d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	byIdSql := "select id, name, city, zipcode, date_of_birth, status from customer where id = ?"

	var c Customer
	err := d.client.Get(&c, byIdSql, id)

	if  err != nil {
		if err == sql.ErrNoRows {
			log.Println("Customer not found " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error when scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}

	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(statusFilter sql.NullString) ([]Customer, *errs.AppError) {


	findAllSql := "select id, name, city, zipcode, date_of_birth, status from customer"

	if statusFilter.Valid {
		findAllSql = findAllSql + " WHERE status = '" + statusFilter.String + "'"
	}

	customers := make([]Customer, 0)
	err :=  d.client.Select(&customers, findAllSql)

	if err != nil {
		log.Println("Error scanning customer query " + err.Error())
		return nil, errs.NewUnexpectedError("Error scanning customer query")
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db := NewSqlxDb()
	return CustomerRepositoryDb{db}
}

