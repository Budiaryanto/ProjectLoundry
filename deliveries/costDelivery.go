package deliveries

import (
	"famLoundry/models"
	"fmt"
)

type CostDelivery struct {
}

func (pd *CostDelivery) PrintOneCost(result *models.Cost) {
	fmt.Printf("%v\n", result)
}
func (pd *CostDelivery) PrintCost(result []*models.Cost) {
	for _, p := range result {
		fmt.Println(p.IdCost, p.ProductName, p.DateOfEntry, p.DateOfEntry, p.Quantity, p.Price, p.Information)
	}
}
func (pd *CostDelivery) PrintTotalCost(result int64) {
	fmt.Printf("Total Cost FamilyLoundry %d\n", result)
}

func NewCostDelivery() *CostDelivery {
	return &CostDelivery{}
}
