package entities

type MonthYear struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Month int    `json:"month"`
	Year  int    `json:"year"`
}
