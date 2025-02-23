package models

type Currency struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"type:varchar(100); not null"`
	Symbol      string `json:"symbol" gorm:"type:varchar(10); not null"`
	Status      bool   `json:"status" gorm:"default:true"`
}
