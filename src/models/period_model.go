package models

import "time"

type Period struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"type:varchar(100); not null"`
	Month       int    `json:"month" gorm:"type:integer; not null"`
	Year        int    `json:"year" gorm:"type:integer; not null"`
	Status      bool   `json:"status" gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
