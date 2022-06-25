package usecase

import (
	"github.com/Rifqi14/omzet_api/domain/model"
	"github.com/Rifqi14/omzet_api/domain/request"
	"github.com/Rifqi14/omzet_api/domain/view_models"
)

type IAuthenticationUseCase interface {
	Login(req *request.LoginRequest) (res view_models.LoginVm, err error)

	GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error)

	GetCurrentUser() (res map[string]interface{}, err error)

	Register(req *request.RegisterRequest) (res model.User, err error)
}
