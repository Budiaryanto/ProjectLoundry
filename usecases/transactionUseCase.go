package usecases

import (
	"famLoundry/models"
	"famLoundry/repositories"
)

type ITransactionUseCase interface {
	RegisterNewTransaction(transaction models.Transaction) (*models.Transaction, error)
	GetTransactionById(id string) (*models.Transaction, error)
	GetTransactionByNoNotaLike(name []string) ([]*models.Transaction, error)
	GetTransactionPaging(pageNo, totalPerPage int) ([]*models.Transaction, error)
	GetTotalTransaction() (int64, error)

	// RegisterNewProductWithPrice(productWithPrice models.ProductWithPrice) (*models.ProductWithPrice, error)
}

type TransactionUseCase struct {
	repo repositories.ITransactionRepository
}

func (p *TransactionUseCase) RegisterNewTransaction(transaction models.Transaction) (*models.Transaction, error) {
	return p.repo.Insert(transaction)
}

func NewTransactionUseCase(repo repositories.ITransactionRepository) ITransactionUseCase {
	return &TransactionUseCase{repo}
}

func (p *TransactionUseCase) GetTransactionById(id string) (*models.Transaction, error) {
	return p.repo.FindOneById(id)
}

func (p *TransactionUseCase) GetTransactionByNoNotaLike(name []string) ([]*models.Transaction, error) {
	result := make([]*models.Transaction, 0)
	for _, q := range name {
		r, err := p.repo.FindAllByNoNotaLike(q)
		if err != nil {
			return nil, err
		}
		result = append(result, r...)
	}
	return result, nil
}

func (p *TransactionUseCase) GetTransactionPaging(pageNo, totalPerPage int) ([]*models.Transaction, error) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if totalPerPage <= 0 {
		totalPerPage = 10
	}
	return p.repo.FindAllTransactionPaging((pageNo-1)*totalPerPage, totalPerPage)
}

func (p *TransactionUseCase) GetTotalTransaction() (int64, error) {
	return p.repo.Count()
}

// func (p *ProductUseCase) RegisterNewProductWithPrice(productWithPrice models.ProductWithPrice) (*models.ProductWithPrice, error) {
// 	return p.repo.InsertProductWithPrice(productWithPrice)
// }
