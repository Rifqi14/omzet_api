package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Rifqi14/omzet_api/package/functioncaller"
	jwtPkg "github.com/Rifqi14/omzet_api/package/jwt"
	"github.com/Rifqi14/omzet_api/package/logruslogger"
	"github.com/Rifqi14/omzet_api/package/messages"
	"github.com/Rifqi14/omzet_api/package/responses"
	"github.com/Rifqi14/omzet_api/package/str"
	"github.com/Rifqi14/omzet_api/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type JwtMiddleware struct {
	*usecase.Contract
}

func (jwtMiddleware JwtMiddleware) New(ctx *fiber.Ctx) (err error) {
	claims := &jwtPkg.CustomClaims{}

	// Check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "middleware-jwt-checkHeader")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	// Check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != t.Method {
			logruslogger.Log(logruslogger.WarnLevel, messages.UnexpectedSigningMethod, functioncaller.PrintFuncName(), "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secret := []byte(jwtMiddleware.JwtCredential.TokenSecret)
		return secret, nil
	})

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "middleware-jwt-checkClaims")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	// Check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		logruslogger.Log(logruslogger.WarnLevel, messages.ExpiredToken, functioncaller.PrintFuncName(), "middleware-jwt-checkTokenLiveTime")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	// Jwe roll back encrypted ID
	jweRes, err := jwtMiddleware.JweCredential.Rollback(claims.Payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwe-rollback")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}
	if jweRes == nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	claims.Id = fmt.Sprintf("%v", jweRes["id"])
	jwtMiddleware.Contract.UserID = uint(str.StringToInt(claims.Id))

	return ctx.Next()
}
