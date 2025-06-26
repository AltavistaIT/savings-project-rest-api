package entities

import "time"

type TableType struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Key         string `json:"key" gorm:"unique; type:varchar(100); not null"`
	Description string `json:"description" gorm:"type:varchar(100); not null"`
	Status      bool   `json:"status" gorm:"default:true"`

	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
