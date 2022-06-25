package command

import "github.com/Rifqi14/omzet_api/domain/model"

type IUserRepository interface {
	Create(user model.User) (res model.User, err error)
}
