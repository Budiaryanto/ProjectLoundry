package models

type Cost struct {
	IdCost      int
	ProductName string
	DateOfEntry string
	Quantity    int64
	Price       int64
	Information string
}

type TotalCost struct {
	Count int64
}
