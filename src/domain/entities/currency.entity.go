package entities

import "time"

type Currency struct {
	ID           uint64 `json:"id" gorm:"primaryKey"`
	Description  string `json:"description" gorm:"type:varchar(100); not null"`
	Abbreviation string `json:"abbreviation" gorm:"type:varchar(10); not null"`
	Symbol       string `json:"symbol" gorm:"type:varchar(10); not null"`
	Status       bool   `json:"status,omitempty" gorm:"default:true"`

	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
}
