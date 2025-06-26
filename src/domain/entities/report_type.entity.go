package entities

import "time"

type ReportType struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"type:varchar(100); not null"`
	Status      bool   `json:"status" gorm:"default:true"`

	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
