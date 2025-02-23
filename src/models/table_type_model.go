package models

type TableType struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"type:varchar(100); not null"`
	Status      bool   `json:"status" gorm:"default:true"`
}
