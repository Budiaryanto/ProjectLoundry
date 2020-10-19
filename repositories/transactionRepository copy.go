package repositories

import (
	"database/sql"
	"errors"
	"famLoundry/models"
	"fmt"
)

type ITransactionRepository interface {
	Insert(transaction models.Transaction) (*models.Transaction, error)
	FindOneById(id string) (*models.Transaction, error)
	FindAllByNoNotaLike(noNota string) ([]*models.Transaction, error)
	FindAllTransactionPaging(pageNo, totalPerPage int) ([]*models.Transaction, error)
	Count() (int64, error)
}

var (
	transactionQueries = map[string]string{
		"transactionFindOneById":              "select customer.idcustomer,name,weight,dateofentry,dateofout,typeofpackage,price,status from customer LEFT JOIN loundrytransaction on customer.idcustomer = loundrytransaction.idcustomer where customer.idcustomer=?",
		"transactionFindAllByNoNotaLike":      "select customer.idcustomer,loundrytransaction.nonota,name,weight,dateofentry,dateofout,typeofpackage,price,status from customer LEFT JOIN loundrytransaction on customer.idcustomer = loundrytransaction.idcustomer where loundrytransaction.nonota like ?",
		"transactionFindAllTransactionPaging": "select customer.idcustomer,loundrytransaction.nonota,name,weight,dateofentry,dateofout,typeofpackage,price,status from customer LEFT JOIN loundrytransaction on customer.idcustomer = loundrytransaction.idcustomer order by nonota limit ?,?",
		"insertTransaction":                   "insert into loundrytransaction(weight,dateofentry,dateofout,typeofpackage,price,status,idcustomer) values(?,?,?,?,?,?,?)",
	}
)

type TransactionRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewTransactionRepository(db *sql.DB) ITransactionRepository {
	ps := make(map[string]*sql.Stmt, len(transactionQueries))
	for n, v := range transactionQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &TransactionRepository{
		db, ps,
	}
}

func (r *TransactionRepository) Insert(transaction models.Transaction) (*models.Transaction, error) {
	idCustomer := transaction.Customer.IdCustomer

	res, err := r.ps["insertTransaction"].Exec(transaction.Weight, transaction.DateOfEntry, transaction.DateOfOut, transaction.TypeOfPackage, transaction.Price, transaction.Status, idCustomer)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindOneById(id string) (*models.Transaction, error) {
	rows, err := r.ps["transactionFindOneById"].Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := new(models.Transaction)
	err = rows.Scan(&res.NoNota, &res.Weight, &res.DateOfEntry, &res.DateOfOut, &res.TypeOfPackage, &res.Price, &res.Status)
	if err != nil {
		panic(err)
	}
	return res, nil
}
func (r *TransactionRepository) FindAllByNoNotaLike(noNota string) ([]*models.Transaction, error) {
	rows, err := r.ps["transactionFindAllByNoNotaLike"].Query(noNota)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Transaction, 0)
	for rows.Next() {
		res := new(models.Transaction)
		err = rows.Scan(&res.NoNota, &res.Weight, &res.DateOfEntry, &res.DateOfOut, &res.TypeOfPackage, &res.Price, &res.Status)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}
func (r *TransactionRepository) FindAllTransactionPaging(pageNo, totalPerPage int) ([]*models.Transaction, error) {
	rows, err := r.ps["transactionFindAllTransactionPaging"].Query(pageNo, totalPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Transaction, 0)
	for rows.Next() {
		res := new(models.Transaction)
		err = rows.Scan(&res.NoNota, &res.Weight, &res.DateOfEntry, &res.DateOfOut, &res.TypeOfPackage, &res.Price, &res.Status)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}

func (r *TransactionRepository) Count() (int64, error) {
	row := r.db.QueryRow("select count(nonota) from transaction")
	res := new(models.TotalTransaction)
	err := row.Scan(&res.Count)
	if err != nil {
		return -1, nil
	}
	return res.Count, nil
}
