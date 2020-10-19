package usecases

import (
	"famLoundry/models"
	"famLoundry/repositories"
)

type ICustomerUseCase interface {
	GetCustomerById(id string) (*models.Customer, error)
	GetCustomerByNameLike(name []string) ([]*models.Customer, error)
	GetCustomerPaging(pageNo, totalPerPage int) ([]*models.Customer, error)
	GetTotalCustomer() (int64, error)
	RegisterNewCustomer(product models.Customer) (*models.Customer, error)
}

type CustomerUseCase struct {
	repo repositories.ICustomerRepository
}

func (p *CustomerUseCase) RegisterNewCustomer(customer models.Customer) (*models.Customer, error) {
	return p.repo.Insert(customer)
}

func NewCustomerUseCase(repo repositories.ICustomerRepository) ICustomerUseCase {
	return &CustomerUseCase{repo}
}

func (p *CustomerUseCase) GetCustomerById(id string) (*models.Customer, error) {
	return p.repo.FindOneById(id)
}

func (p *CustomerUseCase) GetCustomerByNameLike(name []string) ([]*models.Customer, error) {
	result := make([]*models.Customer, 0)
	for _, q := range name {
		r, err := p.repo.FindAllByNameLike(q)
		if err != nil {
			return nil, err
		}
		result = append(result, r...)
	}
	return result, nil
}

func (p *CustomerUseCase) GetCustomerPaging(pageNo, totalPerPage int) ([]*models.Customer, error) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if totalPerPage <= 0 {
		totalPerPage = 10
	}
	return p.repo.FindAllCustomerPaging((pageNo-1)*totalPerPage, totalPerPage)
}
func (p *CustomerUseCase) GetTotalCustomer() (int64, error) {
	return p.repo.Count()
}
