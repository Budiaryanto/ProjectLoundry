package repositories

import (
	"database/sql"
	"errors"
	"famLoundry/models"
	"fmt"
)

type IReportMonthlyRepository interface {
	Insert(reportMonthly models.ReportMonthly) (*models.ReportMonthly, error)
	FindOneById(id int) (*models.ReportMonthly, error)
	ReportMonthlyFindAllByDate(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportMonthly, error)
	FindAllReportMonthlyPaging(dateOfEntryStart string, dateOfEntryFinish string, pageNo int, totalPerPage int) ([]*models.ReportMonthly, error)
	Count() (int64, error)
}

var (
	reportMonthlyQueries = map[string]string{
		"insertReportMonthly":                     "insert into reportmonthly(dateofinput,totalweight, totalincomeinput, totalincomepayment, totalcost, totalactual, information) values(?,?,?,?,?,?,?)",
		"reportMonthlyFindOneById":                "select idreportmonthly,dateofinput,totalweight,totalincomeinput,totalincomepayment,totalcost,totalactual,information from reportmonthly where idreportmonthly=?",
		"reportMonthlyFindAllByDate":              "select idreportmonthly,dateofinput,totalweight,totalincomeinput,totalincomepayment,totalcost,totalactual,information from reportmonthly where dateofinput BETWEEN ? and ?",
		"reportMonthlyFindAllReportMonthlyPaging": "select idreportmonthly,dateofinput,totalweight,totalincomeinput,totalincomepayment,totalcost,totalactual,information from reportmonthly where dateofinput Between ? and ? order by idreportmonthly limit ?,?",
	}
)

type ReportMonthlyRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewReportMonthlyRepository(db *sql.DB) IReportMonthlyRepository {
	ps := make(map[string]*sql.Stmt, len(reportMonthlyQueries))
	for n, v := range reportMonthlyQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &ReportMonthlyRepository{
		db, ps,
	}
}

func (r *ReportMonthlyRepository) Insert(reportMonthly models.ReportMonthly) (*models.ReportMonthly, error) {
	// idReportMonthly := transaction.Customer.IdCustomer

	res, err := r.ps["insertReportMonthly"].Exec(reportMonthly.Date, reportMonthly.TotalWeight, reportMonthly.TotalIncomeInput, reportMonthly.TotalIncomePayment, reportMonthly.TotalCost, reportMonthly.TotalActual, reportMonthly.Information)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &reportMonthly, nil
}

func (r *ReportMonthlyRepository) FindOneById(id int) (*models.ReportMonthly, error) {
	rows := r.ps["reportMonthlyFindOneById"].QueryRow(id)
	res := new(models.ReportMonthly)
	err := rows.Scan(&res.IdReportMonthly, &res.Date, &res.TotalWeight, &res.TotalIncomeInput, &res.TotalIncomePayment, &res.TotalCost, &res.TotalActual, &res.Information)
	if err != nil {
		panic(err)
	}
	return res, nil
}

func (r *ReportMonthlyRepository) ReportMonthlyFindAllByDate(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportMonthly, error) {
	rows, err := r.ps["reportMonthlyFindAllByDate"].Query(dateOfEntryStart, dateOfEntryFinish)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.ReportMonthly, 0)
	for rows.Next() {
		res := new(models.ReportMonthly)
		err = rows.Scan(&res.IdReportMonthly, &res.Date, &res.TotalWeight, &res.TotalIncomeInput, &res.TotalIncomePayment, &res.TotalCost, &res.TotalActual, &res.Information)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}
func (r *ReportMonthlyRepository) FindAllReportMonthlyPaging(dateOfEntryStart string, dateOfEntryFinish string, pageNo int, totalPerPage int) ([]*models.ReportMonthly, error) {
	rows, err := r.ps["reportMonthlyFindAllReportMonthlyPaging"].Query(dateOfEntryStart, dateOfEntryFinish, pageNo, totalPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.ReportMonthly, 0)
	for rows.Next() {
		res := new(models.ReportMonthly)
		err = rows.Scan(&res.IdReportMonthly, &res.Date, &res.TotalWeight, &res.TotalIncomeInput, &res.TotalIncomePayment, &res.TotalCost, &res.TotalActual, &res.Information)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil

}

func (r *ReportMonthlyRepository) Count() (int64, error) {
	row := r.db.QueryRow("select count(idreportmonthly) from reportmonthly")
	res := new(models.TotalReportMonthly)
	err := row.Scan(&res.Count)
	if err != nil {
		return -1, nil
	}
	return res.Count, nil
}
