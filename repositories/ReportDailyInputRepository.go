package repositories

import (
	"database/sql"
	"famLoundry/models"
)

type IReportDailyInputRepository interface {
	// Insert(transaction models.Transaction) (*models.Transaction, error)
	// FindOneById(id string) (*models.Transaction, error)
	ReportDailyInputFindAllByDate(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportDailyInput, error)
	FindAllReportDailyInputPaging(pageNo, totalPerPage int) ([]*models.ReportDailyInput, error)
}

var (
	reportDailyInputQueries = map[string]string{
		"reportDailyInputFindAllByDate": "select customer.idcustomer,name,weight,dateofentry,dateofout, typeofpackage,price,status from customer LEFT JOIN loundrytransaction on loundrytransaction.idcustomer = customer.idcustomer where loundrytransaction.dateofentry BETWEEN ? and ?",
		// "transactionFindAllByNoNotaLike":      "select customer.idcustomer,transaction.nonota,name,weight,dateofentry,dateofout,typeofpackage,price,status from customer LEFT JOIN transaction on customer.idcustomer = transaction.idcustomer where transaction.nonota like ?",
		"reportDailyInputFindAllReportDailyInputPaging": "select customer.idcustomer,loundrytransaction.nonota,name,weight,dateofentry,dateofout,typeofpackage,price,status from customer LEFT JOIN loundrytransaction on customer.idcustomer = loundrytransaction.idcustomer order by nonota limit ?,?",
		// "insertReportDailyInput":              "insert into reportdailyinput(idcustomer,nonota,dateofentry,information) values(?,?,?,?)",
	}
)

type ReportDailyInputRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewReportDailyInputRepository(db *sql.DB) IReportDailyInputRepository {
	ps := make(map[string]*sql.Stmt, len(reportDailyInputQueries))
	for n, v := range reportDailyInputQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &ReportDailyInputRepository{
		db, ps,
	}
}

func (r *ReportDailyInputRepository) ReportDailyInputFindAllByDate(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportDailyInput, error) {
	rows, err := r.ps["reportDailyInputFindAllByDate"].Query(dateOfEntryStart, dateOfEntryFinish)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.ReportDailyInput, 0)
	for rows.Next() {
		res := new(models.ReportDailyInput)
		err = rows.Scan(&res.Customer.IdCustomer, &res.Customer.Name, &res.Transaction.Weight, &res.Transaction.DateOfEntry, &res.Transaction.DateOfOut, &res.Transaction.TypeOfPackage, &res.Transaction.Price, &res.Transaction.Status)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}
func (r *ReportDailyInputRepository) FindAllReportDailyInputPaging(pageNo, totalPerPage int) ([]*models.ReportDailyInput, error) {
	rows, err := r.ps["reportDailyInputFindAllReportDailyInputPaging"].Query(pageNo, totalPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.ReportDailyInput, 0)
	for rows.Next() {
		res := new(models.ReportDailyInput)
		err = rows.Scan(&res.Customer.IdCustomer, &res.Customer.Name, &res.Transaction.Weight, &res.Transaction.DateOfEntry, &res.Transaction.DateOfOut, &res.Transaction.TypeOfPackage, &res.Transaction.Price, &res.Transaction.Status, &res.Transaction.DateOfEntry)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}
