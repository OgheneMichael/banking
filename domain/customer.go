package domain

import "banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	// FindAll status can be 1 -for active- or 0 -for inactive- or can be empty -for all-
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
