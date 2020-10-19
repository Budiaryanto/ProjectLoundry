package models

type Customer struct {
	IdCustomer string
	Name       string
	Address    string
	Contact    string
}

type TotalCustomer struct {
	Count int64
}
