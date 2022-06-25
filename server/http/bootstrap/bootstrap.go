package bootstrap

import (
	"github.com/Rifqi14/omzet_api/usecase"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Bootstrap struct {
	App        *fiber.App
	DB         *gorm.DB
	UcContract usecase.Contract
	Validator  *validator.Validate
	Translator ut.Translator
}
