package deliveries

import (
	"famLoundry/models"
	"fmt"
)

type TransactionDelivery struct {
}

func (pd *TransactionDelivery) PrintOneTransaction(result *models.Transaction) {
	fmt.Printf("%v\n", result)
}
func (pd *TransactionDelivery) PrintTransaction(result []*models.Transaction) {
	for _, p := range result {
		fmt.Println(p.NoNota, p.Customer.IdCustomer, p.Weight, p.DateOfEntry, p.DateOfOut, p.TypeOfPackage, p.Price, p.Status)
	}
}
func (pd *TransactionDelivery) PrintTotalTransaction(result int64) {
	fmt.Printf("Total Transaction FamilyLoundry %d\n", result)
}

func NewTransactionDelivery() *TransactionDelivery {
	return &TransactionDelivery{}
}
