package usecase

import (
	"github.com/Rifqi14/omzet_api/domain/request"
	"github.com/Rifqi14/omzet_api/domain/view_models"
)

type IOmzetUseCase interface {
	ReportByMerchant(merchantId int, req request.ReportMerchantRequest) (res []view_models.ReportMerchantVm, pagination view_models.PaginationVm, err error)

	ReportByOutlet(outletId int, req request.ReportMerchantRequest) (res []view_models.ReportOutletVm, pagination view_models.PaginationVm, err error)
}
