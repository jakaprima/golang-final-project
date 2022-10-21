package controllers

import (
	"encoding/json"
	"finalproject/models"
	"finalproject/views"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateResource godoc
// @Summary Post SocialMedia
// @Description Post SocialMedia
// @Tags SOCIAL_MEDIA
// @Accept  json
// @Produce      json
// @Param Social_Media body views.Swagger_Social_Media_Post true "Post Social Media"
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Success 200  {object} string "success"
// @Router /socialmedias [post]
func (postgres *HandlersController) SocialMedias_Post(ctx *gin.Context) {
	// Check Authorization
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "request does not contain an access token.",
			},
		})
		ctx.Abort()
		return
	}
	if strings.Contains(tokenString, "Bearer") == false {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "format Authentification Bearer not found",
			},
		})
		ctx.Abort()
		return
	}
	jwtString := strings.Split(tokenString, "Bearer ")[1]
	err1 := ValidateToken(tokenString)
	if err1 != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: err1.Error(),
			},
		})
		ctx.Abort()
		return
	}
	ctx.Next()

	// decode/Extract JWT
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Verivication"), nil
	})
	username := fmt.Sprintf("%v", claims["username"])
	email := fmt.Sprintf("%v", claims["email"])
	println(username)
	println(email)

	// Get Body Value
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	var key_data views.Request_Social_Medias

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "failed to post photo",
			},
		})
		return
	}
	println("Name: ", key_data.Name)
	println("Request_Social_Medias: ", key_data.Social_Media_Url)

	// Tittle and Photo_Url Validation
	if key_data.Name == "" || key_data.Social_Media_Url == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "name or social media url field is empty!",
			},
		})
		return
	}

	// query data from table photo
	var result models.User
	postgres.db.Table("users").Select("id").Where("email = ?", email).Scan(&result)
	userid := int(result.ID)

	if result.ID == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "Token is not valid!",
			},
		})
		return
	}

	//generate Social_Media_ID
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	println(userid)
	userid_creat := userid
	// Create/Insert Photo to databases
	err_photo := postgres.db.Create(&models.SocialMedia{
		ID:                r1.Int() / 100000,
		Name:              key_data.Name,
		Social_Media_Url:  key_data.Social_Media_Url,
		Profile_Image_Url: key_data.Profile_Image_Url,
		User_Id:           userid_creat,
		Create_At:         time.Now(),
		Update_At:         time.Now(),
	}).Error

	if err_photo != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "failed to post photo to db",
			},
		})
		return
	}

	WriteJsonResponse_PostSocialMedia(ctx, &views.Post_Social_Media{
		Message_Action: "SUCCESS",
		Status:         http.StatusCreated,
		Message_Data: views.Post_Social_Media_Data{
			ID:               r1.Int() / 100000,
			Name:             key_data.Name,
			Social_Media_Url: key_data.Social_Media_Url,
			User_Id:          userid,
			Create_At:        time.Now(),
		},
	})
}

// CreateResource godoc
// @Summary Get SocialMedia
// @Description Get Social_Media
// @Tags SOCIAL_MEDIA
// @Accept  json
// @Produce      json
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Success 200  {object} string "success"
// @Router /socialmedias [get]
func (postgres *HandlersController) SocialMedias_Get(ctx *gin.Context) {
	// Check Authorization
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "request does not contain an access token.",
			},
		})
		ctx.Abort()
		return
	}
	if strings.Contains(tokenString, "Bearer") == false {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "format Authentification Bearer not found",
			},
		})
		ctx.Abort()
		return
	}
	jwtString := strings.Split(tokenString, "Bearer ")[1]
	err1 := ValidateToken(tokenString)
	if err1 != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: err1.Error(),
			},
		})
		ctx.Abort()
		return
	}
	ctx.Next()

	// decode/Extract JWT
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Verivication"), nil
	})
	username := fmt.Sprintf("%v", claims["username"])
	email := fmt.Sprintf("%v", claims["email"])
	println(username)
	println(email)

	var result_id models.User
	postgres.db.Table("users").Select("id").Where("email = ?", email).Scan(&result_id)
	userid := int(result_id.ID)

	var result []models.SocialMedia
	err := postgres.db.Table("social_media").Where("user_id = ?", userid).First(&result).Error
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "No Social Media Found!!",
			},
		})
		return
	}
	hasil := postgres.db.Table("social_media").Where("user_id = ?", userid).Find(&result)
	if result[0].ID == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "No Social Media Found!!",
			},
		})
		return
	}
	var a = make([]interface{}, hasil.RowsAffected)
	for i := 0; i < int(hasil.RowsAffected); i++ {
		result_social_media := &views.User_Sosial_Media_Data_get{
			ID:               result[i].ID,
			Name:             result[i].Name,
			Social_Media_Url: result[i].Social_Media_Url,
			User_Id:          userid,
			Create_At:        result[i].Create_At,
			Update_At:        result[i].Update_At,
			User_Get: views.User_SM_Get{
				ID:                userid,
				Username:          username,
				Profile_Image_Url: "di isi manual",
			},
		}
		a[i] = result_social_media
	}
	WriteJsonResponse_GetSocialMedia(ctx, &views.Post_Social_Media_Get{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data: views.Post_Social_Media_Get_Data{
			Social_Media: a,
		},
	})
}

// CreateResource godoc
// @Summary Update Social_Media
// @Description Update Social_Media
// @Tags SOCIAL_MEDIA
// @Accept  json
// @Produce      json
// @Param Social_Media body views.Swagger_Social_Media_Put true "Update Social Media"
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Param socialMediaId path int true "Social Media ID"
// @Success 200  {object} string "success"
// @Router /socialmedias/{socialMediaId} [put]
func (postgres *HandlersController) SocialMedias_Put(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	println(socialMediaId)
	// Check Authorization
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "request does not contain an access token.",
			},
		})
		ctx.Abort()
		return
	}
	if strings.Contains(tokenString, "Bearer") == false {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "format Authentification Bearer not found",
			},
		})
		ctx.Abort()
		return
	}
	jwtString := strings.Split(tokenString, "Bearer ")[1]
	err1 := ValidateToken(tokenString)
	if err1 != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: err1.Error(),
			},
		})
		ctx.Abort()
		return
	}
	ctx.Next()

	// decode/Extract JWT
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Verivication"), nil
	})
	username := fmt.Sprintf("%v", claims["username"])
	email := fmt.Sprintf("%v", claims["email"])
	println(username)
	println(email)

	// Get Body Value
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	var key_data views.Request_Social_Medias

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "failed to update Social Media Data!!",
			},
		})
		return
	}
	println("Name: %s", key_data.Name)
	println("Social_Media_Url: %s", key_data.Social_Media_Url)
	println("Profile_Image_Url: %s", key_data.Profile_Image_Url)

	// querry user data
	var result models.User
	postgres.db.Table("users").Select("id").Where("email = ?", email).Scan(&result)
	println(result.ID)

	//Query Social Media ID
	var result_sm models.SocialMedia
	postgres.db.Table("social_media").Where("id = ?", socialMediaId).Find(&result_sm)
	if result_sm.ID == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "No Social Media Found on this Account!",
			},
		})
		return
	}

	//// Update Data
	postgres.db.Model(&models.SocialMedia{}).Where("id = ?", result_sm.ID).Updates(
		models.SocialMedia{
			Name:             key_data.Name,
			Social_Media_Url: key_data.Social_Media_Url,
			User_Id:          result.ID,
			Update_At:        time.Now(),
		})

	WriteJsonResponse_PutSocialMedia(ctx, &views.Put_Social_Media_Get{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data: views.Put_Social_Media_Data{
			ID:               result_sm.ID,
			Name:             key_data.Name,
			Social_Media_Url: key_data.Social_Media_Url,
			User_Id:          result.ID,
			Update_At:        time.Now(),
		},
	})
}

// CreateResource godoc
// @Summary Get Social_Media
// @Description Get Social_MEdia
// @Tags SOCIAL_MEDIA
// @Accept  json
// @Produce      json
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Param socialMediaId path int true "Social Media ID"
// @Success 200  {object} string "success"
// @Router /socialmedias/{socialMediaId} [delete]
func (postgres *HandlersController) SocialMedias_Delete(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	println(socialMediaId)
	// Check Authorization
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "request does not contain an access token.",
			},
		})
		ctx.Abort()
		return
	}
	if strings.Contains(tokenString, "Bearer") == false {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "format Authentification Bearer not found",
			},
		})
		ctx.Abort()
		return
	}
	jwtString := strings.Split(tokenString, "Bearer ")[1]
	err1 := ValidateToken(tokenString)
	if err1 != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: err1.Error(),
			},
		})
		ctx.Abort()
		return
	}
	ctx.Next()

	// decode/Extract JWT
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Verivication"), nil
	})
	username := fmt.Sprintf("%v", claims["username"])
	email := fmt.Sprintf("%v", claims["email"])
	println(username)
	println(email)

	// querry user data
	var result models.User
	postgres.db.Table("users").Select("id").Where("email = ?", email).Scan(&result)
	println(result.ID)

	//Query Social Media ID
	var result_sm models.SocialMedia
	postgres.db.Table("social_media").Where("id = ? AND user_id = ?", socialMediaId, result.ID).Find(&result_sm)
	if result_sm.ID == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "The Social Media ID Found on this Account!",
			},
		})
		return
	}

	err := postgres.db.Where("id = ?", socialMediaId).Delete(&models.SocialMedia{}).Error
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Delete User Failed",
			},
		})
		return
	}

	WriteJsonResponse_Delete(ctx, &views.Resp_Delete{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data: views.Data_Delete{
			Message: "Your social media has been successfully deleted!!",
		},
	})

}
