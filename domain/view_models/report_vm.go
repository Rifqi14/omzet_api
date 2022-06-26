package view_models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Rifqi14/omzet_api/domain/model"
)

type ReportMerchantVm struct {
	Date         string  `json:"date"`
	MerchantName string  `json:"merchant_name"`
	Omzet        float64 `json:"omzet"`
}

type ReportVm struct {
	Merchant []ReportMerchantVm `json:"merchant"`
}

func NewReportVm() *ReportVm {
	return &ReportVm{}
}

func (vm ReportVm) BuildMerchant(merchants []model.Transaction, days, month, year, limit, offset, page int) (res []ReportMerchantVm) {
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
			fmt.Println("Date: ", merchant.CreatedAt.Format("2006-01-02"), "Merchant: ", merchant.Merchant.MerchantName, "Total: ", total)
		}
		view.Omzet = total
		res = append(res, view)
	}

	return res
}
