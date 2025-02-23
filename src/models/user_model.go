package models

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey "`
	Username string `json:"username" gorm:"unique; type varchar(100); not null"`
	Email    string `json:"email" gorm:"unique; type varchar(100); not null"`
	Password string `json:"password" gorm:"not null"`
	Status   bool   `json:"status" gorm:"default:true"`
}
