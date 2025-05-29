package entities

import "time"

type TransactionType struct {
	ID          uint64     `json:"id" gorm:"primaryKey"`
	Key         string     `json:"key" gorm:"unique; type:varchar(100); not null"`
	Description string     `json:"description" gorm:"type:varchar(100); not null"`
	Status      bool       `json:"status" gorm:"default:true"`
	TableTypeID uint64     `json:"table_type_id" gorm:"not null"`
	TableType   *TableType `json:"table_type,omitempty" gorm:"foreignKey:TableTypeID"`

	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
