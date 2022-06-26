package query

import "github.com/Rifqi14/omzet_api/domain/model"

type IOmzetRepository interface {
	ReportByMerchant(merchantId int, search, orderBy, sort, start_period, end_period string, limit, offset int64) (res []model.Transaction, count int64, err error)

	ReportByOutlet(outletId int, search, orderBy, sort, start_period, end_period string, limit, offset int64) (res []model.Transaction, count int64, err error)
}
