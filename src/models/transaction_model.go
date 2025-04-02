package models

import "time"

type Transaction struct {
	ID              uint64           `json:"id" gorm:"primaryKey"`
	Description     string           `json:"description" gorm:"type:varchar(100); not null"`
	TypeID          uint64           `json:"type_id" gorm:"not null"`
	TransactionType *TransactionType `json:"transaction_type,omitempty" gorm:"foreignKey:TypeID; constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Budget          float64          `json:"budget" gorm:"type:decimal(10,2); not null"`
	CurrencyID      uint64           `json:"currency_id" gorm:"not null"`
	Currency        *Currency        `json:"currency,omitempty" gorm:"foreignKey:CurrencyID; constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Status          bool             `json:"status" gorm:"default:true"`

	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
