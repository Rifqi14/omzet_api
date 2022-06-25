package query

import (
	"github.com/Rifqi14/omzet_api/domain/model"
	"github.com/Rifqi14/omzet_api/domain/repository/query"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewQueryUserRepository(db *gorm.DB) query.IUserRepository {
	return UserRepository{
		DB: db,
	}
}

func (repo UserRepository) ReadBy(column, operator string, value interface{}) (res []model.User, err error) {
	tx := repo.DB

	err = tx.Find(&res, column+" "+operator+" ?", value).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo UserRepository) First(column, operator string, value interface{}) (res model.User, err error) {
	tx := repo.DB

	err = tx.First(&res, column+" "+operator+" ?", value).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
