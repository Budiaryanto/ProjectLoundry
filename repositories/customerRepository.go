package repositories

import (
	"database/sql"
	"errors"
	"famLoundry/models"
	"fmt"
)

type ICustomerRepository interface {
	Insert(customer models.Customer) (*models.Customer, error)
	FindOneById(id string) (*models.Customer, error)
	FindAllByNameLike(name string) ([]*models.Customer, error)
	FindAllCustomerPaging(pageNo, totalPerPage int) ([]*models.Customer, error)
	Count() (int64, error)
}

var (
	customerQueries = map[string]string{
		"customerFindOneById":           "select idcustomer,name,address,contact from customer where idcustomer=?",
		"customerFindAllByNameLike":     "select idcustomer,name,address,contact from customer where name like ?",
		"customerFindAllCustomerPaging": "select idcustomer,name,address,contact from customer order by idcustomer limit ?,?",
		"insertCustomer":                "insert into customer(idcustomer,name,address,contact) values(?,?,?,?)",
	}
)

type CustomerRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewCustomerRepository(db *sql.DB) ICustomerRepository {
	ps := make(map[string]*sql.Stmt, len(customerQueries))
	for n, v := range customerQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &CustomerRepository{
		db, ps,
	}
}

func (r *CustomerRepository) Insert(customer models.Customer) (*models.Customer, error) {
	// id := guuid.New()
	// customer.IdCustomer = id.String()
	res, err := r.ps["insertCustomer"].Exec(customer.IdCustomer, customer.Name, customer.Address, customer.Contact)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &customer, nil
}

func (r *CustomerRepository) FindOneById(id string) (*models.Customer, error) {
	rows, err := r.ps["customerFindOneById"].Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := new(models.Customer)
	err = rows.Scan(&res.IdCustomer, &res.Name, &res.Address, &res.Contact)
	if err != nil {
		panic(err)
	}
	return res, nil
}
func (r *CustomerRepository) FindAllByNameLike(name string) ([]*models.Customer, error) {
	rows, err := r.ps["customerFindAllByNameLike"].Query(name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Customer, 0)
	for rows.Next() {
		res := new(models.Customer)
		err = rows.Scan(&res.IdCustomer, &res.Name, &res.Address, &res.Contact)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}
func (r *CustomerRepository) FindAllCustomerPaging(pageNo, totalPerPage int) ([]*models.Customer, error) {
	rows, err := r.ps["customerFindAllCustomerPaging"].Query(pageNo, totalPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Customer, 0)
	for rows.Next() {
		res := new(models.Customer)
		err = rows.Scan(&res.IdCustomer, &res.Name, &res.Address, &res.Contact)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}

func (r *CustomerRepository) Count() (int64, error) {
	row := r.db.QueryRow("select count(idcustomer) from customer")
	res := new(models.TotalCustomer)
	err := row.Scan(&res.Count)
	if err != nil {
		return -1, nil
	}
	return res.Count, nil
}
