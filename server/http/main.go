package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Rifqi14/omzet_api/domain"
	"github.com/Rifqi14/omzet_api/package/functioncaller"
	"github.com/Rifqi14/omzet_api/package/logruslogger"
	"github.com/Rifqi14/omzet_api/server/http/bootstrap"
	"github.com/Rifqi14/omzet_api/usecase"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

var (
	logFormat = `{"host":"${host}","pid":"${pid}","time":"${time}","request-id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","in":"${bytesReceived}","out":"${bytesSent}"}`
	validatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-env")
	}

	config, err := domain.LoadConfig()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-config")
	}

	dbConn := config.DB
	gormDb, err := config.DB.DB()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "connection-error")
	}
	defer gormDb.Close()

	app := fiber.New()

	ValidatorInit()

	ucContract := usecase.Contract{
		App:           app,
		DB:            dbConn,
		JweCredential: config.JweCredential,
		JwtCredential: config.JwtCredential,
		Validate:      validatorDriver,
		Redis:         config.Redis,
		Translator:    translator,
	}

	// Bootstrap init
	boot := bootstrap.Bootstrap{
		App:        app,
		DB:         dbConn,
		UcContract: ucContract,
		Validator:  validatorDriver,
		Translator: translator,
	}

	boot.App.Use(recover.New())
	boot.App.Use(requestid.New())
	boot.App.Use(cors.New())
	boot.App.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	boot.AppRoute()
	log.Fatal(boot.App.Listen(os.Getenv("APP_HOST")))
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	validatorDriver = validator.New()

	err := enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	if err != nil {
		fmt.Println(err)
	}

	err = idTranslations.RegisterDefaultTranslations(validatorDriver, transID)
	if err != nil {
		fmt.Println(err)
	}
	switch os.Getenv("APP_LOCALE") {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
