package view_models

import (
	"errors"
	"strconv"
	"time"

	"github.com/Rifqi14/omzet_api/domain/model"
	"github.com/Rifqi14/omzet_api/package/messages"
)

type ReportMerchantVm struct {
	Date         string  `json:"date"`
	MerchantName string  `json:"merchant_name"`
	Omzet        float64 `json:"omzet"`
}

type ReportOutletVm struct {
	Date         string  `json:"date"`
	MerchantName string  `json:"merchant_name"`
	OutletName   string  `json:"outlet_name"`
	Omzet        float64 `json:"omzet"`
}

type ReportVm struct {
	Merchant []ReportMerchantVm `json:"merchant"`
	Outlet   []ReportOutletVm   `json:"outlet"`
}

func NewReportVm() *ReportVm {
	return &ReportVm{}
}

func (vm ReportVm) BuildMerchant(merchants []model.Transaction, days, month, year, limit, offset, page int, userId uint) (res []ReportMerchantVm, err error) {
	for i := offset + 1; i <= limit*page; i++ {
		day := strconv.Itoa(i)
		if i < 10 {
			day = "0" + day
		}
		date, _ := time.Parse("2006-01-02", strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-"+day)
		view := ReportMerchantVm{}
		view.Date = date.Format("2006-01-02")
		total := 0.0
		for _, merchant := range merchants {

			view.MerchantName = merchant.Merchant.MerchantName
			if merchant.CreatedAt.Format("2006-01-02") == date.Format("2006-01-02") {
				total += merchant.BillTotal
			}
			if merchant.Merchant.UserID != userId {
				return nil, errors.New(messages.Unauthorized)
			}
		}
		view.Omzet = total
		res = append(res, view)
	}

	return res, nil
}

func (vm ReportVm) BuildOutlet(outlets []model.Transaction, days, month, year, limit, offset, page int, userId uint) (res []ReportOutletVm, err error) {
	for i := offset + 1; i <= limit*page; i++ {
		day := strconv.Itoa(i)
		if i < 10 {
			day = "0" + day
		}
		date, _ := time.Parse("2006-01-02", strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-"+day)
		view := ReportOutletVm{}
		view.Date = date.Format("2006-01-02")
		total := 0.0
		for _, outlet := range outlets {

			view.MerchantName = outlet.Outlet.Merchant.MerchantName
			view.OutletName = outlet.Outlet.OutletName
			if outlet.CreatedAt.Format("2006-01-02") == date.Format("2006-01-02") {
				total += outlet.BillTotal
			}
			if outlet.Outlet.Merchant.UserID != userId {
				return nil, errors.New(messages.Unauthorized)
			}
		}
		view.Omzet = total
		res = append(res, view)
	}

	return res, nil
}
