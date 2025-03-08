package models

import "time"

type Report struct {
	ID          uint64     `json:"id" gorm:"primaryKey"`
	Description string     `json:"description" gorm:"type:varchar(100); not null"`
	TypeID      uint64     `json:"type_id" gorm:"not null"`
	ReportType  ReportType `json:"report_type" gorm:"foreignKey:TypeID"`
	UserID      uint64     `json:"user_id" gorm:"not null"`
	User        User       `json:"user" gorm:"foreignKey:UserID"`
	PeriodID    uint64     `json:"period_id" gorm:"not null"`
	Period      Period     `json:"period" gorm:"foreignKey:PeriodID"`
	Content     string     `json:"content" gorm:"type:json; not null"`
	Status      bool       `json:"status" gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
