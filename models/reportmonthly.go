package models

type ReportMonthly struct {
	IdReportMonthly    int64
	Date               string
	TotalWeight        float64
	TotalIncomeInput   int64
	TotalIncomePayment int64
	TotalCost          int64
	TotalActual        int64
	Information        string
}

type TotalReportMonthly struct {
	Count int64
}
