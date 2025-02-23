package models

type ReportType struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"type:varchar(100); not null"`
	Status      bool   `json:"status" gorm:"default:true"`
}
