package deliveries

import (
	"famLoundry/models"
	"fmt"
)

type CustomerDelivery struct {
}

func (pd *CustomerDelivery) PrintOneCustomer(result *models.Customer) {
	fmt.Printf("%v\n", result)
}
func (pd *CustomerDelivery) PrintCustomer(result []*models.Customer) {
	for _, p := range result {
		fmt.Println(p.IdCustomer, p.Name, p.Address, p.Contact)
	}
}
func (pd *CustomerDelivery) PrintTotalCustomer(result int64) {
	fmt.Printf("Total Customer FamilyLoundry %d\n", result)
}

func NewCustomerDelivery() *CustomerDelivery {
	return &CustomerDelivery{}
}
