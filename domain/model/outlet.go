package model

import "time"

type Outlet struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MerchantID uint      `json:"merchant_id"`
	OutletName string    `json:"outlet_name"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
	CreatedBy  uint      `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:current_timestamp;not null"`
	UpdatedBy  uint      `json:"updated_by"`
}
