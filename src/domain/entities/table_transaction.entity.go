package entities

import "time"

type TableTransaction struct {
	ID            uint64       `json:"id" gorm:"primaryKey"`
	TableID       uint64       `json:"table_id" gorm:"not null"`
	Table         *Table       `json:"table,omitempty" gorm:"foreignKey:TableID"`
	TransactionID uint64       `json:"transaction_id" gorm:"not null"`
	Transaction   *Transaction `json:"transaction,omitempty" gorm:"foreignKey:TransactionID"`
	Position      int          `json:"position" gorm:"type:integer; not null"`

	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
