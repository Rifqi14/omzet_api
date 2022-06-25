package model

import "time"

type Transaction struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MerchantID uint      `json:"merchant_id"`
	OutletID   uint      `json:"outlet_id"`
	BillTotal  float64   `json:"bill_total"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
	CreatedBy  uint      `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:current_timestamp;not null"`
	UpdatedBy  uint      `json:"updated_by"`
}
