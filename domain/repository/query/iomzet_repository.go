package query

import "github.com/Rifqi14/omzet_api/domain/model"

type IOmzetRepository interface {
	ReportByMerchant(merchantId int) (res []model.Transaction, err error)

	ReportByOutlet(outletId int) (res []model.Transaction, err error)
}
