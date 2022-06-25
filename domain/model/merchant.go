package model

import "time"

type Merchant struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       uint      `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
	CreatedBy    uint      `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:current_timestamp;not null"`
	UpdatedBy    uint      `json:"updated_by"`
}
