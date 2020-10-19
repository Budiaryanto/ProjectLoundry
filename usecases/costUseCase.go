package usecases

import (
	"famLoundry/models"
	"famLoundry/repositories"
)

type ICostUseCase interface {
	RegisterNewCost(product models.Cost) (*models.Cost, error)
	GetCostById(idCost int64) (*models.Cost, error)
	GetCostByProductNameLike(productName []string) ([]*models.Cost, error)
	GetCostByDateOfEntryLike(dateOfEntry []string) ([]*models.Cost, error)
	GetCostPaging(pageNo, totalPerPage int) ([]*models.Cost, error)
	GetTotalCost() (int64, error)
}

type CostUseCase struct {
	repo repositories.ICostRepository
}

func NewCostUseCase(repo repositories.ICostRepository) ICostUseCase {
	return &CostUseCase{repo}
}

func (p *CostUseCase) RegisterNewCost(cost models.Cost) (*models.Cost, error) {
	return p.repo.Insert(cost)
}

func (p *CostUseCase) GetCostById(idCost int64) (*models.Cost, error) {
	return p.repo.FindOneById(idCost)
}

func (p *CostUseCase) GetCostByProductNameLike(productName []string) ([]*models.Cost, error) {
	result := make([]*models.Cost, 0)
	for _, q := range productName {
		r, err := p.repo.FindAllByProductNameLike(q)
		if err != nil {
			return nil, err
		}
		result = append(result, r...)
	}
	return result, nil
}

func (p *CostUseCase) GetCostByDateOfEntryLike(dateOfEntry []string) ([]*models.Cost, error) {
	result := make([]*models.Cost, 0)
	for _, q := range dateOfEntry {
		r, err := p.repo.FindAllByDateOfEntryLike(q)
		if err != nil {
			return nil, err
		}
		result = append(result, r...)
	}
	return result, nil
}

func (p *CostUseCase) GetCostPaging(pageNo, totalPerPage int) ([]*models.Cost, error) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if totalPerPage <= 0 {
		totalPerPage = 10
	}
	return p.repo.FindAllCostPaging((pageNo-1)*totalPerPage, totalPerPage)
}
func (p *CostUseCase) GetTotalCost() (int64, error) {
	return p.repo.Count()
}
