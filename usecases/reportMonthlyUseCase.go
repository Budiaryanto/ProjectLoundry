package usecases

import (
	"famLoundry/models"
	"famLoundry/repositories"
)

type IReportMonthlyUseCase interface {
	RegisterNewReportMonthly(reportMonthly models.ReportMonthly) (*models.ReportMonthly, error)
	GetReportMonthlyById(id int) (*models.ReportMonthly, error)
	GetAllReportMonthlyByDateLike(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportMonthly, error)
	GetReportMonthlyPaging(dateOfEntryStart string, dateOfEntryFinish string, pageNo int, totalPerPage int) ([]*models.ReportMonthly, error)
	GetTotalReportMonthly() (int64, error)
}

type ReportMonthlyUseCase struct {
	repo repositories.IReportMonthlyRepository
}

func NewReportMonthlyUseCase(repo repositories.IReportMonthlyRepository) IReportMonthlyUseCase {
	return &ReportMonthlyUseCase{repo}
}

func (p *ReportMonthlyUseCase) RegisterNewReportMonthly(reportMonthly models.ReportMonthly) (*models.ReportMonthly, error) {
	return p.repo.Insert(reportMonthly)
}

func (p *ReportMonthlyUseCase) GetReportMonthlyById(id int) (*models.ReportMonthly, error) {
	return p.repo.FindOneById(id)
}

func (p *ReportMonthlyUseCase) GetAllReportMonthlyByDateLike(dateOfEntryStart, dateOfEntryFinish string) ([]*models.ReportMonthly, error) {
	return p.repo.ReportMonthlyFindAllByDate(dateOfEntryStart, dateOfEntryFinish)
}

func (p *ReportMonthlyUseCase) GetReportMonthlyPaging(dateOfEntryStart string, dateOfEntryFinish string, pageNo int, totalPerPage int) ([]*models.ReportMonthly, error) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if totalPerPage <= 0 {
		totalPerPage = 10
	}
	return p.repo.FindAllReportMonthlyPaging(dateOfEntryStart, dateOfEntryFinish, (pageNo-1)*totalPerPage, totalPerPage)
}

func (p *ReportMonthlyUseCase) GetTotalReportMonthly() (int64, error) {
	return p.repo.Count()
}
