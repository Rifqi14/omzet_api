package v1

import (
	"time"

	"github.com/Rifqi14/omzet_api/domain/request"
	ucinterface "github.com/Rifqi14/omzet_api/domain/usecase"
	"github.com/Rifqi14/omzet_api/domain/view_models"
	"github.com/Rifqi14/omzet_api/package/functioncaller"
	"github.com/Rifqi14/omzet_api/package/logruslogger"
	"github.com/Rifqi14/omzet_api/repositories/query"
	"github.com/Rifqi14/omzet_api/usecase"
)

type OmzetUseCase struct {
	*usecase.Contract
}

func NewOmzetUseCase(contract *usecase.Contract) ucinterface.IOmzetUseCase {
	return &OmzetUseCase{contract}
}

func (uc OmzetUseCase) ReportByMerchant(merchantId int, req request.ReportMerchantRequest) (res []view_models.ReportMerchantVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repo := query.NewQueryOmzetRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	reports, _, err := repo.ReportByMerchant(merchantId, req.Search, orderBy, sort, req.StartPeriod, req.EndPeriod, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-report-merchant")
		return nil, pagination, err
	}

	date, err := time.Parse("2006-01-02", req.EndPeriod)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-end-period")
		return nil, pagination, err
	}
	res = view_models.NewReportVm().BuildMerchant(reports, date.Day(), int(date.Month()), int(date.Year()), int(limit), int(offset), int(page))

	pagination = uc.SetPaginationResponse(page, limit, int64(date.Day()))

	return res, pagination, nil
}
