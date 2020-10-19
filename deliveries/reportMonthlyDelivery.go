package deliveries

import (
	"famLoundry/models"
	"fmt"
)

type ReportMonthlyDelivery struct {
}

func (pd *ReportMonthlyDelivery) PrintOneReportMonthly(result *models.ReportMonthly) {
	fmt.Printf("%v\n", result)
}
func (pd *ReportMonthlyDelivery) PrintReportMonthly(result []*models.ReportMonthly) {
	for _, p := range result {
		fmt.Println(p.IdReportMonthly, p.Date, p.TotalWeight, p.TotalIncomeInput, p.TotalIncomePayment, p.TotalCost, p.TotalActual, p.Information)
	}
}
func (pd *ReportMonthlyDelivery) PrintReportMonthlyPage(result []*models.ReportMonthly) {
	for _, p := range result {
		fmt.Println(p.IdReportMonthly, p.Date, p.TotalWeight, p.TotalIncomeInput, p.TotalIncomePayment, p.TotalCost, p.TotalActual, p.Information)
	}
}
func (pd *ReportMonthlyDelivery) PrintTotalReportMonthly(result int64) {
	fmt.Printf("Total Report Monthly FamilyLoundry %d\n", result)
}

func NewReportMonthlyDelivery() *ReportMonthlyDelivery {
	return &ReportMonthlyDelivery{}
}
