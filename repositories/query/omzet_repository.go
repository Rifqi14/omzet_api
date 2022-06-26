package query

import (
	"github.com/Rifqi14/omzet_api/domain/model"
	"github.com/Rifqi14/omzet_api/domain/repository/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OmzetRepository struct {
	DB *gorm.DB
}

func NewQueryOmzetRepository(db *gorm.DB) query.IOmzetRepository {
	return OmzetRepository{
		DB: db,
	}
}

func (repo OmzetRepository) ReportByMerchant(merchantId int, search, orderBy, sort, start_period, end_period string, limit, offset int64) (res []model.Transaction, count int64, err error) {
	tx := repo.DB

	err = tx.Debug().Preload(clause.Associations).Preload("Merchant.User").Model(&model.Transaction{}).Count(&count).Where("merchant_id = ?", merchantId).Where("created_at >= ?", start_period).Where("created_at <= ?", end_period).Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (repo OmzetRepository) ReportByOutlet(outletId int, search, orderBy, sort, start_period, end_period string, limit, offset int64) (res []model.Transaction, count int64, err error) {
	tx := repo.DB

	err = tx.Debug().Preload(clause.Associations).Preload("Outlet.Merchant.User").Model(&model.Transaction{}).Count(&count).Where("outlet_id = ?", outletId).Where("created_at >= ?", start_period).Where("created_at <= ?", end_period).Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}
