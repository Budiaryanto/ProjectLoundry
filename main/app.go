package main

import (
	"database/sql"
	"famLoundry/config"
	"famLoundry/deliveries"
	"famLoundry/repositories"
	"famLoundry/usecases"

	_ "github.com/go-sql-driver/mysql"
)

type app struct {
	db *sql.DB
}

func newApp() app {
	c := config.NewConfig()
	err := c.InitDb()
	if err != nil {
		panic(err)
	}
	myapp := app{
		db: c.Db,
	}
	return myapp
}

func (a app) run() {
	repo := repositories.NewReportMonthlyRepository(a.db)
	usecase := usecases.NewReportMonthlyUseCase(repo)
	delivery := deliveries.NewReportMonthlyDelivery()

	//#INSERT
	// _, err := usecase.RegisterNewReportMonthly(models.ReportMonthly{
	// 	Date:               "2020-10-18",
	// 	TotalWeight:        112.45,
	// 	TotalIncomeInput:   1050000,
	// 	TotalIncomePayment: 950000,
	// 	TotalCost:          245000,
	// 	TotalActual:        874000,
	// 	Information:        "Bala-Bala",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	//#FindById
	// idReportMonthly := 3
	// result, err := usecase.GetReportMonthlyById(idReportMonthly)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)

	//#FindByDate
	// dateOfEntryStart := "2020-10-16"
	// dateOfEntryFinish := "2020-10-17"
	// res, err := (usecase.GetAllReportMonthlyByDateLike(dateOfEntryStart, dateOfEntryFinish))
	// delivery.PrintReportMonthly(res)
	// if err != nil {
	// 	panic(err)
	// }

	//#FindByPagging
	dateOfEntryStart := "2020-10-16"
	dateOfEntryFinish := "2020-10-18"
	pageNo := 2
	totalPage := 4
	result, err := usecase.GetReportMonthlyPaging(dateOfEntryStart, dateOfEntryFinish, pageNo, totalPage)
	delivery.PrintReportMonthlyPage(result)

	if err != nil {
		panic(err)
	}

	//#GetTotal
	// dateOfEntryStart := "2020-10-16"
	// dateOfEntryFinish := "2020-10-20"
	// result, err := usecase.GetAllReportMonthlyByDateLike(dateOfEntryStart, dateOfEntryFinish)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result[0])
}
func main() {
	newApp().run()
}
