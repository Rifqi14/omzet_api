package v1

import (
	"errors"
	"os"
	"strconv"

	"github.com/Rifqi14/omzet_api/domain/model"
	"github.com/Rifqi14/omzet_api/domain/request"
	iusecase "github.com/Rifqi14/omzet_api/domain/usecase"
	"github.com/Rifqi14/omzet_api/domain/view_models"
	"github.com/Rifqi14/omzet_api/package/functioncaller"
	"github.com/Rifqi14/omzet_api/package/hashing"
	"github.com/Rifqi14/omzet_api/package/logruslogger"
	"github.com/Rifqi14/omzet_api/package/messages"
	"github.com/Rifqi14/omzet_api/repositories/query"
	"github.com/Rifqi14/omzet_api/usecase"
)

type AuthenticationUseCase struct {
	*usecase.Contract
}

func NewAuthenticationUseCase(contract *usecase.Contract) iusecase.IAuthenticationUseCase {
	return &AuthenticationUseCase{Contract: contract}
}

func (uc AuthenticationUseCase) Login(req *request.LoginRequest) (res view_models.LoginVm, err error) {
	db := uc.DB
	repo := query.NewQueryUserRepository(db)
	user, err := repo.First("user_name", "=", req.UserName)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.UserNotFound, functioncaller.PrintFuncName(), "uc-check-user-name")
		return res, err
	}

	if isPasswordValid := hashing.CheckHashString(req.Password, user.Password); !isPasswordValid {
		logruslogger.Log(logruslogger.WarnLevel, messages.CredentialDoNotMatch, functioncaller.PrintFuncName(), "uc-check-password")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	// Generate jwt payload and encrypted with jwe
	payload := map[string]interface{}{
		"id":        user.ID,
		"user_name": user.UserName,
		"name":      user.Name,
	}
	jwePayload, err := uc.JweCredential.GenerateJwePayload(payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "Error generate jwe payload", functioncaller.PrintFuncName(), "uc-generate-jwe-payload")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	// Generate jwt token
	res, err = uc.GenerateJWT(req.UserName, jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "Error generate jwt token", functioncaller.PrintFuncName(), "uc-generate-jwt-token")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	// Store token to redis
	err = uc.Redis.StoreToRedisWithExpired("token-"+strconv.Itoa(int(user.ID)), res, os.Getenv("TOKEN_EXP_TIME")+"m")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "Error store token to redis", functioncaller.PrintFuncName(), "uc-store-token-to-redis")
	}

	// Store data user to redis
	userLoggedIn := map[string]interface{}{
		"id":        user.ID,
		"user_name": user.UserName,
		"name":      user.Name,
	}
	err = uc.Redis.StoreToRedisWithExpired("user-"+strconv.Itoa(int(user.ID)), userLoggedIn, os.Getenv("TOKEN_EXP_TIME")+"m")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "Error store user to redis", functioncaller.PrintFuncName(), "uc-store-user-to-redis")
	}

	return res, nil
}

func (uc AuthenticationUseCase) GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error) {
	res.Token, res.TokenExpiration, err = uc.JwtCredential.GetToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "error get token", functioncaller.PrintFuncName(), "uc-generate-jwt")
		return res, err
	}

	res.RefreshToken, res.RefreshTokenExpiration, err = uc.JwtCredential.GetRefreshToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "error get refresh token", functioncaller.PrintFuncName(), "uc-generate-jwt")
		return res, err
	}
	return res, nil
}

func (uc AuthenticationUseCase) GetCurrentUser() (res map[string]interface{}, err error) {
	err = uc.Redis.GetFromRedis("user-"+strconv.Itoa(int(uc.UserID)), &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "error get user from redis", functioncaller.PrintFuncName(), "uc-get-current-user")
		return nil, err
	}

	return res, nil
}

func (uc AuthenticationUseCase) Register(req *request.RegisterRequest) (res model.User, err error) {
	panic("Implement me")
}
