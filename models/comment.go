package models

import (
	"time"
)

type Comment struct {
	ID        int    `gorm:"primaryKey"`
	User_Id   int    `json:"user_id"`
	Photo_Id  int    `json:"photo_id"`
	Message   string `json:"message"`
	Create_At time.Time
	Update_At time.Time
}
