package models

type ReportDailyInput struct {
	IdReportDaylyInput int64
	Customer           Customer
	Transaction        Transaction
	StartDate          string
	FinishDate         string
	Information        string
}

type TotalDaily struct {
	Count int64
}
