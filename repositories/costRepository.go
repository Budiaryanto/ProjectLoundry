package repositories

import (
	"database/sql"
	"errors"
	"famLoundry/models"
	"fmt"
)

type ICostRepository interface {
	Insert(cost models.Cost) (*models.Cost, error)
	FindOneById(idCost int64) (*models.Cost, error)
	FindAllByProductNameLike(productName string) ([]*models.Cost, error)
	FindAllByDateOfEntryLike(dateOfEntry string) ([]*models.Cost, error)
	FindAllCostPaging(pageNo, totalPerPage int) ([]*models.Cost, error)
	Count() (int64, error)
}

var (
	costQueries = map[string]string{
		"costFindOneById":              "select idcost,productname,dateofentry,quantity,price,information from cost where idcost=?",
		"costFindAllByProductNameLike": "select idcost,productname,dateofentry,quantity,price,information from cost where productname like ?",
		"costFindAllByDateOfEntryLike": "select idcost,productname,dateofentry,quantity,price,information from cost where dateofentry like ?",
		"FindAllCostPaging":            "select idcost,productname,dateofentry,quantity,price,information from cost order by idcost limit ?,?",
		"insertCost":                   "insert into cost(productname,dateofentry,quantity,price,information) values(?,?,?,?,?)",
	}
)

type CostRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewCostRepository(db *sql.DB) ICostRepository {
	ps := make(map[string]*sql.Stmt, len(costQueries))
	for n, v := range costQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &CostRepository{
		db, ps,
	}
}

func (r *CostRepository) Insert(cost models.Cost) (*models.Cost, error) {
	// id := guuid.New()
	// customer.IdCustomer = id.String()
	res, err := r.ps["insertCost"].Exec(cost.ProductName, cost.DateOfEntry, cost.Quantity, cost.Price, cost.Information)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &cost, nil
}

func (r *CostRepository) FindOneById(idCost int64) (*models.Cost, error) {
	rows := r.ps["costFindOneById"].QueryRow(idCost)

	res := new(models.Cost)
	err := rows.Scan(&res.IdCost, &res.ProductName, &res.DateOfEntry, &res.Quantity, &res.Price, &res.Information)
	fmt.Println(err)
	if err != sql.ErrNoRows && err != nil {
		panic(err)
	}

	return res, nil
	// rows, err := r.ps["costFindOneById"].Query(idCost)
	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()
	// result := make([]*models.Cost, 0)
	// for rows.Next() {
	// 	res := new(models.Cost)
	// 	err = rows.Scan(&res.IdCost, &res.ProductName, &res.DateOfEntry, &res.Quantity, &res.Price, &res.Information)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	result = append(result, res)
	// }
	// return result[0], nil
}

func (r *CostRepository) FindAllByProductNameLike(productName string) ([]*models.Cost, error) {
	rows, err := r.ps["costFindAllByProductNameLike"].Query(productName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Cost, 0)
	for rows.Next() {
		res := new(models.Cost)
		err = rows.Scan(&res.IdCost, &res.ProductName, &res.DateOfEntry, &res.Quantity, &res.Price, &res.Information)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}

func (r *CostRepository) FindAllByDateOfEntryLike(dateOfEntry string) ([]*models.Cost, error) {
	rows, err := r.ps["costFindAllByDateOfEntryLike"].Query(dateOfEntry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Cost, 0)
	for rows.Next() {
		res := new(models.Cost)
		err = rows.Scan(&res.IdCost, &res.ProductName, &res.DateOfEntry, &res.Quantity, &res.Price, &res.Information)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}

func (r *CostRepository) FindAllCostPaging(pageNo, totalPerPage int) ([]*models.Cost, error) {
	rows, err := r.ps["FindAllCostPaging"].Query(pageNo, totalPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Cost, 0)
	for rows.Next() {
		res := new(models.Cost)
		err = rows.Scan(&res.IdCost, &res.ProductName, &res.DateOfEntry, &res.Quantity, &res.Price, &res.Information)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}

func (r *CostRepository) Count() (int64, error) {
	row := r.db.QueryRow("select count(idcost) from cost")
	res := new(models.TotalCost)
	err := row.Scan(&res.Count)
	if err != nil {
		return -1, nil
	}
	return res.Count, nil
}
