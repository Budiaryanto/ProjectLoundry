package usecases

import (
	"famLoundry/models"
	"famLoundry/repositories"
)

type IReportDailyInputUseCase interface {
	// RegisterNewTransaction(transaction models.Transaction) (*models.Transaction, error)
	// GetTransactionById(id string) (*models.Transaction, error)
	GetAllReportDailyInputByDateOfEntryLike(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportDailyInput, error)
	GetReportDailyInputPaging(pageNo, totalPerPage int) ([]*models.ReportDailyInput, error)
	// GetTotalTransaction() (int64, error)

	// RegisterNewProductWithPrice(productWithPrice models.ProductWithPrice) (*models.ProductWithPrice, error)
}

type ReportDailyInputUseCase struct {
	repo repositories.IReportDailyInputRepository
}

func NewReportDailyInputUseCase(repo repositories.IReportDailyInputRepository) IReportDailyInputUseCase {
	return &ReportDailyInputUseCase{repo}
}

func (p *ReportDailyInputUseCase) GetAllReportDailyInputByDateOfEntryLike(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportDailyInput, error) {
	return p.repo.ReportDailyInputFindAllByDate(dateOfEntryStart, dateOfEntryFinish)
}

func (p *ReportDailyInputUseCase) GetReportDailyInputPaging(pageNo, totalPerPage int) ([]*models.ReportDailyInput, error) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if totalPerPage <= 0 {
		totalPerPage = 10
	}
	return p.repo.FindAllReportDailyInputPaging((pageNo-1)*totalPerPage, totalPerPage)
}
