package v1

import (
	"strconv"

	ihandlers "github.com/Rifqi14/omzet_api/domain/handlers"
	"github.com/Rifqi14/omzet_api/domain/request"
	"github.com/Rifqi14/omzet_api/package/messages"
	"github.com/Rifqi14/omzet_api/package/responses"
	"github.com/Rifqi14/omzet_api/server/http/handlers"
	v1 "github.com/Rifqi14/omzet_api/usecase/v1"
	"github.com/gofiber/fiber/v2"
)

type OmzetHandler struct {
	handlers.Handler
}

func NewOmzetHandler(handler handlers.Handler) ihandlers.IOmzetHandler {
	return &OmzetHandler{
		Handler: handler,
	}
}

func (h OmzetHandler) GetReportMerchant(ctx *fiber.Ctx) (err error) {
	req := new(request.ReportMerchantRequest)
	if err := ctx.QueryParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	merchantId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	// Service processing
	uc := v1.NewOmzetUseCase(h.UcContract)
	data, meta, err := uc.ReportByMerchant(merchantId, *req)
	if err != nil && err.Error() == messages.Unauthorized {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusUnauthorized, messages.Unauthorized, err)).Send(ctx)
	}
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusInternalServerError, messages.FailedLoadPayload, err)).Send(ctx)
	}

	// Response
	return responses.NewResponse(responses.ResponseSuccess(data, meta, "Success")).Send(ctx)
}

func (h OmzetHandler) GetReportOutlet(ctx *fiber.Ctx) (err error) {
	req := new(request.ReportMerchantRequest)
	if err := ctx.QueryParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	outlet_id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	// Service processing
	uc := v1.NewOmzetUseCase(h.UcContract)
	data, meta, err := uc.ReportByOutlet(outlet_id, *req)
	if err != nil && err.Error() == messages.Unauthorized {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusUnauthorized, messages.Unauthorized, err)).Send(ctx)
	}
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusInternalServerError, messages.FailedLoadPayload, err)).Send(ctx)
	}

	// Response
	return responses.NewResponse(responses.ResponseSuccess(data, meta, "Success")).Send(ctx)
}
