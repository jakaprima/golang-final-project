package views

import (
	"time"
)

// Success Response
type Response struct {
	Message string      `json:"message" example:"GET_SUCCESS"`
	Status  int         `json:"status" example:"201"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Failed Response
type Message struct {
	Message string `json:"message"`
}
type Failed struct {
	Message_Action string  `json:"message_action" example:"GET_SUCCESS"`
	Status         int     `json:"status" example:"201"`
	Message_Data   Message `json:"message_data"`
}

// /////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////// Success Response For USER Table ////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////
type Data_Register struct {
	Age       int    `json:"age"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Update_At time.Time
	Create_At time.Time
}
type Resp_Register_Success struct {
	Message_Action string        `json:"message_action" example:"GET_SUCCESS"`
	Status         int           `json:"status" example:"201"`
	Message_Data   Data_Register `json:"message_data"`
}

type Token struct {
	Token string `json:"token"`
}
type Resp_Login struct {
	Message_Action string `json:"message_action" example:"GET_SUCCESS"`
	Status         int    `json:"status" example:"201"`
	Message_Data   Token  `json:"message_data"`
}

type Put struct {
	Age       int    `json:"age"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Update_At time.Time
}
type Resp_Put struct {
	Message_Action string `json:"message_action" example:"GET_SUCCESS"`
	Status         int    `json:"status" example:"201"`
	Message_Data   Put    `json:"message_data"`
}

type Data_Delete struct {
	Message string `json:"message"`
}
type Resp_Delete struct {
	Message_Action string      `json:"message_action" example:"GET_SUCCESS"`
	Status         int         `json:"status" example:"201"`
	Message_Data   Data_Delete `json:"message_data"`
}

// /////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////// Success Response For PHOTO Table ////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////
type Data_Photo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
	Create_At time.Time
}
type Resp_Post_Photo struct {
	Message_Action string     `json:"message_action" example:"GET_SUCCESS"`
	Status         int        `json:"status" example:"201"`
	Message_Data   Data_Photo `json:"message_data"`
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Get_Photos_Data struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
	Create_At time.Time
	User      User `json:"user"`
}

type Get_Photos struct {
	Message_Action string      `json:"message_action" example:"GET_SUCCESS"`
	Status         int         `json:"status" example:"201"`
	Message_Data   interface{} `json:"message_data,omitempty"`
}

type Put_Photos_Data struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
	Update_At time.Time
}

type Put_Photos struct {
	Message_Action string          `json:"message_action" example:"GET_SUCCESS"`
	Status         int             `json:"status" example:"201"`
	Message_Data   Put_Photos_Data `json:"message_data"`
}

// /////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////// Success Response For COMMENT Table ////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////
type Comments_Post_Data struct {
	ID        int    `gorm:"primaryKey"`
	Message   string `json:"message"`
	Photo_Id  int    `json:"photo_id" gorm:"foreignKey:PhotoRefer"`
	User_Id   int    `json:"user_id" gorm:"foreignKey:UserRefer"`
	Create_At time.Time
}
type Comments_Post struct {
	Message_Action string             `json:"message_action" example:"GET_SUCCESS"`
	Status         int                `json:"status" example:"201"`
	Message_Data   Comments_Post_Data `json:"message_data"`
}

//// GET Fuction coment

type Coment_User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Coment_Photo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
}

type Comments_Put_Data struct {
	ID        int    `gorm:"primaryKey"`
	Message   string `json:"message"`
	Photo_Id  int    `json:"photo_id" gorm:"foreignKey:PhotoRefer"`
	User_Id   int    `json:"user_id" gorm:"foreignKey:UserRefer"`
	Create_At time.Time
	Update_at time.Time
	User      Coment_User  `json:"User"`
	Photo     Coment_Photo `json:"Photo"`
}

type Comment_Get struct {
	Message_Action string      `json:"message_action" example:"GET_SUCCESS"`
	Status         int         `json:"status" example:"201"`
	Message_Data   interface{} `json:"message_data,omitempty"`
}

type Get_Comment struct {
	Message_Action string      `json:"message_action" example:"GET_SUCCESS"`
	Status         int         `json:"status" example:"201"`
	Message_Data   interface{} `json:"message_data,omitempty"`
}

type Put_Comment_Data struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
	Update_At time.Time
}

type Put_Comment struct {
	Message_Action string           `json:"message_action" example:"GET_SUCCESS"`
	Status         int              `json:"status" example:"201"`
	Message_Data   Put_Comment_Data `json:"message_data"`
}

// /////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////// Success Response For Social Media Table ////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////

type Post_Social_Media_Data struct {
	ID               int    `gorm:"primaryKey"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id          int    `json:"user_id" gorm:"foreignKey:UserRefer"`
	Create_At        time.Time
}

type Post_Social_Media struct {
	Message_Action string                 `json:"message_action" example:"GET_SUCCESS"`
	Status         int                    `json:"status" example:"201"`
	Message_Data   Post_Social_Media_Data `json:"message_data"`
}

type User_SM_Get struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	Profile_Image_Url string `json:"profile_image_url"`
}

type User_Sosial_Media_Data_get struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id          int    `json:"user_id" gorm:"foreignKey:UserRefer"`
	Create_At        time.Time
	Update_At        time.Time
	User_Get         User_SM_Get `json:"User"`
}

type Post_Social_Media_Get_Data struct {
	Social_Media interface{} `json:"social_media"`
}

type Post_Social_Media_Get struct {
	Message_Action string                     `json:"message_action" example:"GET_SUCCESS"`
	Status         int                        `json:"status" example:"201"`
	Message_Data   Post_Social_Media_Get_Data `json:"message_data"`
}

type Put_Social_Media_Data struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_Id          int    `json:"user_id" gorm:"foreignKey:UserRefer"`
	Update_At        time.Time
}

type Put_Social_Media_Get struct {
	Message_Action string                `json:"message_action" example:"GET_SUCCESS"`
	Status         int                   `json:"status" example:"201"`
	Message_Data   Put_Social_Media_Data `json:"message_data"`
}
