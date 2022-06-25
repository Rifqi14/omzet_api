package query

import "github.com/Rifqi14/omzet_api/domain/model"

type IUserRepository interface {
	ReadBy(column, operator string, value interface{}) (res []model.User, err error)

	First(column, operator string, value interface{}) (res model.User, err error)
}
