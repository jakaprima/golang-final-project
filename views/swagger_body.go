package views

// / USer Swagger ///
type Swagger_User_Register_Post struct {
	Username string `json:"username" example:"String"`
	Email    string `json:"email" example:"String"`
	Password string `json:"password" example:"String"`
	Age      int    `json:"age" example:"0"`
}

type Swagger_User_Login_Post struct {
	Email    string `json:"email" example:"String"`
	Password string `json:"password" example:"String"`
}

type Swagger_User_Update_Put struct {
	Email    string `json:"email" example:"String"`
	Username string `json:"username" example:"String"`
}

// / Photo Swagger //////
type Swagger_Photo_Register_Post struct {
	Title     string `json:"title" example:"String"`
	Caption   string `json:"caption" example:"String"`
	Photo_Url string `json:"photo_url" example:"String"`
}

type Swagger_Photo_Register_Put struct {
	Title     string `json:"title" example:"String"`
	Caption   string `json:"caption" example:"String"`
	Photo_Url string `json:"photo_url" example:"String"`
}

// Comments Swagger /////
type Swagger_Comment_Register_Post struct {
	Message  string `json:"message" example:"String"`
	Photo_Id int    `json:"photo_id" example:"0"`
}

type Swagger_Comment_Register_Put struct {
	Message string `json:"message" example:"String"`
}

// Social Media Swagger ////
type Swagger_Social_Media_Post struct {
	Name              string `json:"name" example:"String"`
	Social_Media_Url  string `json:"social_media_url" example:"String"`
	Profile_Image_Url string `json:"profile_image_url" example:"String"`
}

type Swagger_Social_Media_Put struct {
	Name             string `json:"name" example:"String"`
	Social_Media_Url string `json:"social_media_url" example:"String"`
}
