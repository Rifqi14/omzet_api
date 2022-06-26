package request

type ReportMerchantRequest struct {
	Search      string `json:"search" form:"search" query:"search"`
	OrderBy     string `json:"order_by" form:"order_by" query:"order_by"`
	Sort        string `json:"sort" form:"sort" query:"sort"`
	StartPeriod string `json:"start_period" form:"start_period" query:"start_period"`
	EndPeriod   string `json:"end_period" form:"end_period" query:"end_period"`
	Limit       int64  `json:"limit" form:"limit" query:"limit"`
	Offset      int64  `json:"offset" form:"offset" query:"offset"`
}
