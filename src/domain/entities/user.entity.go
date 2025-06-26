package entities

import "time"

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey "`
	Username string `json:"username" gorm:"unique; type varchar(100); not null"`
	Email    string `json:"email" gorm:"unique; type varchar(100); not null"`
	Password string `json:"password" gorm:"not null"`
	Status   bool   `json:"status" gorm:"default:true"`

	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
