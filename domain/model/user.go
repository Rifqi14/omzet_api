package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
	CreatedBy uint      `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp;not null"`
	UpdatedBy uint      `json:"updated_by"`
}
