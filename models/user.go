package models

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age" gorm:"not null"`
	Create_At time.Time
	Update_At time.Time
}
