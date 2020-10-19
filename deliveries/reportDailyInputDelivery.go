package deliveries

import (
	"famLoundry/models"
	"fmt"
)

type ReportDailyInputDelivery struct {
}

func (pd *ReportDailyInputDelivery) PrintOneReportDailyInput(result *models.ReportDailyInput) {
	fmt.Printf("%v\n", result)
}
func (pd *ReportDailyInputDelivery) PrintReportDailyInput(result []*models.ReportDailyInput) {
	for _, p := range result {
		fmt.Println(p.Customer.IdCustomer, p.Customer.Name, p.Transaction.Weight, p.Transaction.TypeOfPackage, p.Transaction.Price, p.Transaction.Status, p.Transaction.DateOfEntry)
	}
}
func (pd *ReportDailyInputDelivery) PrintTotalReportDailyInput(result int64) {
	fmt.Printf("Total Report Daily Input FamilyLoundry %d\n", result)
}

func NewReportDailyInputDelivery() *ReportDailyInputDelivery {
	return &ReportDailyInputDelivery{}
}
