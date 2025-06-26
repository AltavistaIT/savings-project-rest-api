package entities

import "time"

type Table struct {
	ID        uint64     `json:"id" gorm:"primaryKey"`
	UserID    uint64     `json:"user_id" gorm:"not null"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	TypeID    uint64     `json:"type_id" gorm:"not null"`
	TableType *TableType `json:"table_type,omitempty" gorm:"foreignKey:TypeID"`
	Amount    float64    `json:"amount" gorm:"type:decimal(10,2); not null; default:0"`
	Status    bool       `json:"status" gorm:"default:true"`
	MonthYear string     `json:"month_year" gorm:"type:text; not null; check: month_year ~ '^[0-9]{4}-[0-9]{2}$'"`

	TableTransactions []*TableTransaction `json:"table_transactions,omitempty" gorm:"foreignKey:TableID"`

	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
