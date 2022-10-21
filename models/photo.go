package models

import (
	"time"
)

type Photo struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
	Create_At time.Time
	Update_At time.Time
}