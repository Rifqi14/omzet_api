package domain

import (
	"fmt"
	"log"
	"os"

	"github.com/Rifqi14/omzet_api/package/jwe"
	"github.com/Rifqi14/omzet_api/package/jwt"
	"github.com/Rifqi14/omzet_api/package/orm"
	redisPkg "github.com/Rifqi14/omzet_api/package/redis"
	"github.com/Rifqi14/omzet_api/package/str"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v7"
	jwtFiber "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
)

type Config struct {
	DB            *gorm.DB
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	JwtConfig     jwtFiber.Config
	Validator     *validator.Validate
	Redis         redisPkg.RedisClient
}

var (
	ValidatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	Translator      ut.Translator
)

func LoadConfig() (res Config, err error) {
	// Setup DB Connection

	dbInfo := orm.Connection{
		Host:                    os.Getenv("DB_HOST"),
		Port:                    os.Getenv("DB_PORT"),
		User:                    os.Getenv("DB_USER"),
		Password:                os.Getenv("DB_PASS"),
		DbName:                  os.Getenv("DB_NAME"),
		DBMaxConnection:         str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		DBMAxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		DBMaxLifeTimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFE_TIME_CONNECTION")),
	}
	res.DB, err = dbInfo.Conn()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Jwe Credential
	res.JweCredential = jwe.Credential{
		KeyLocation: os.Getenv("JWE_PRIVATE_KEY"),
		Passphrase:  os.Getenv("JWE_PRIVATE_KEY_PASSPHRASE"),
	}

	// Jwt Credential
	res.JwtCredential = jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	// Jwt Config
	res.JwtConfig = jwtFiber.Config{
		SigningKey: []byte(res.JwtCredential.TokenSecret),
		Claims:     &jwt.CustomClaims{},
	}

	// Setup redis
	redisOption := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	res.Redis = redisPkg.RedisClient{Client: redis.NewClient(redisOption)}
	pong, err := res.Redis.Client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Redis ping status: "+pong, err)

	res.Validator = ValidatorDriver

	return res, err
}
