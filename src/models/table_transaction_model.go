package models

import "time"

type TableTransaction struct {
	ID            uint64      `json:"id" gorm:"primaryKey"`
	TableID       uint64      `json:"table_id" gorm:"not null"`
	Table         Table       `json:"table" gorm:"foreignKey:TableID"`
	TransactionID uint64      `json:"transaction_id" gorm:"not null"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
