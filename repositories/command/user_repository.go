package command

import (
	"github.com/Rifqi14/omzet_api/domain/model"
	"github.com/Rifqi14/omzet_api/domain/repository/command"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewCommandUserRepository(db *gorm.DB) command.IUserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) Create(user model.User) (res model.User, err error) {
	panic("implement me")
}
