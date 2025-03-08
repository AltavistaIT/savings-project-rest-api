package models

import "time"

type Table struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	UserID    uint64    `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	TypeID    uint64    `json:"type_id" gorm:"not null"`
	TableType TableType `json:"table_type" gorm:"foreignKey:TypeID"`
	PeriodID  uint64    `json:"period_id" gorm:"not null"`
	Period    Period    `json:"period" gorm:"foreignKey:PeriodID"`
	Amount    float64   `json:"amount" gorm:"type:decimal(10,2); not null"`
	Status    bool      `json:"status" gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
