package models

type Transaction struct {
	ID                uint64          `json:"id" gorm:"primaryKey"`
	Description       string          `json:"description" gorm:"type:varchar(100); not null"`
	TransactionTypeID uint64          `json:"transaction_type_id" gorm:"not null"`
	TransactionType   TransactionType `gorm:"foreignKey:TransactionTypeID; constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Budget            float64         `json:"budget" gorm:"type:decimal(10,2); not null"`
	CurrentBudget     float64         `json:"current_budget" gorm:"type:decimal(10,2); not null"`
	CurrencyID        uint64          `json:"currency_id" gorm:"not null"`
	Currency          Currency        `gorm:"foreignKey:CurrencyID; constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Status            bool            `json:"status" gorm:"default:true"`
}
