package models

import (
	"time"
)

type SocialMedia struct {
	ID                int    `gorm:"primaryKey"`
	Name              string `json:"name"`
	Social_Media_Url  string `json:"social_media_url"`
	Profile_Image_Url string `json:"profile_image_url"`
	User_Id           int    `json:"user_id"`
	Create_At         time.Time
	Update_At         time.Time
}
