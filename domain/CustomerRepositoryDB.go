package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/nicholasanthonys/hexagonal-banking/errs"
	"github.com/nicholasanthonys/hexagonal-banking/logger"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, err

	}

	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {
		logger.Error("Error while scanning customers " + err.Error())
		return nil, err

	}

	return customers, nil

}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Unexpected database error")
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	return &c, nil

}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{dbClient}
}
