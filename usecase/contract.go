package usecase

import (
	"github.com/Rifqi14/omzet_api/domain/view_models"
	"github.com/Rifqi14/omzet_api/package/jwe"
	"github.com/Rifqi14/omzet_api/package/jwt"
	"github.com/Rifqi14/omzet_api/package/redis"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Contract struct {
	UserID        uint
	App           *fiber.App
	DB            *gorm.DB
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	Validate      *validator.Validate
	Redis         redis.RedisClient
	Translator    ut.Translator
}

const (
	// Default limit for pagination
	defaultLimit = 15

	// Default order by
	defaultOrderBy = "created_at"

	// Default sort
	defaultSort = "desc"

	// Default last page for pagination
	defaultLastPage = 1
)

func (uc Contract) SetPaginationParameter(page, limit int64, order, sort string) (int64, int64, int64, string, string) {
	if page <= 0 {
		page = 1
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) SetPaginationResponse(page, limit, total int64) (res view_models.PaginationVm) {
	var lastPage int64

	if total > 0 {
		lastPage = total / limit
		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}
