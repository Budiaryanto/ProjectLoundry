package models

type Transaction struct {
	NoNota        string
	Customer      Customer
	Weight        float64
	DateOfEntry   string
	DateOfOut     string
	TypeOfPackage string
	Price         int64
	Status        string
}

type TotalTransaction struct {
	Count int64
}
