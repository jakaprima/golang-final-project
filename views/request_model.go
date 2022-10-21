package views

import "time"

type Request_Register struct {
	Id_Number int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	Create_At time.Time
	Update_At time.Time
}

type Request_Photos struct {
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url`
}

type Request_Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Request_Update struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Request_Coments struct {
	Message  string `json:"message"`
	Photo_Id int    `json:"photo_id"`
}

type Request_Social_Medias struct {
	Name              string `json:"name"`
	Social_Media_Url  string `json:"social_media_url"`
	Profile_Image_Url string `json:"profile_image_url"`
}
