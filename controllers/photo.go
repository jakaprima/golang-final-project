package controllers

import (
	"database/sql"
	"encoding/json"
	"finalproject/models"
	"finalproject/views"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateResource godoc
// @Summary Post Photo
// @Description Post Photo
// @Tags PHOTO
// @Accept  json
// @Produce      json
// @Param Photo body views.Swagger_Photo_Register_Post true "Post Photo"
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Success 200  {object} string "success"
// @Router /photos [post]
func (postgres *HandlersController) Post_Photos(ctx *gin.Context) {
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

	var key_data views.Request_Photos

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
	println("Title: %s", key_data.Title)
	println("Caption: %s", key_data.Caption)
	println("Photo_Url: %s", key_data.Photo_Url)

	// Tittle and Photo_Url Validation
	if key_data.Title == "" || key_data.Photo_Url == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "photo_url or title field is empty!",
			},
		})
		return
	}

	// query data from table photo
	var result models.User
	postgres.db.Table("users").Select("id").Where("email = ?", email).Scan(&result)
	println(result.ID)
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

	//generate photo_ID
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Create/Insert Photo to databases
	err_photo := postgres.db.Create(&models.Photo{
		ID:        r1.Int() / 100000,
		Title:     key_data.Title,
		Caption:   key_data.Caption,
		Photo_Url: key_data.Photo_Url,
		User_Id:   result.ID,
		Create_At: time.Now(),
		Update_At: time.Now(),
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

	WriteJsonResponse_PostPhoto(ctx, &views.Resp_Post_Photo{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data: views.Data_Photo{
			ID:        r1.Int() / 100000,
			Title:     key_data.Title,
			Caption:   key_data.Caption,
			Photo_Url: key_data.Photo_Url,
			User_Id:   result.ID,
			Create_At: time.Now(),
		},
	})
}

// CreateResource godoc
// @Summary Get Photo
// @Description Get Photo
// @Tags PHOTO
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Success 201  {object} string "success"
// @Router /photos [get]
func (postgres *HandlersController) Get_Photos(ctx *gin.Context) {
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

	var s sql.NullString
	postgres.db.Select("id").Where("email = ?", email).First(&models.User{}).Scan(&s)
	user_id := s.String
	fmt.Printf("%s \n", user_id)
	if user_id == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "No Photo Found!!",
			},
		})
		return
	}

	var result []models.Photo
	err := postgres.db.Table("photos").Where("user_id = ?", user_id).First(&models.Photo{}).Error
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "No Photo Found!!",
			},
		})
		return
	}
	hasil := postgres.db.Table("photos").Where("user_id = ?", user_id).Find(&result)

	var a = make([]interface{}, hasil.RowsAffected)
	for i := 0; i < int(hasil.RowsAffected); i++ {
		tambah := &views.Get_Photos_Data{
			ID:        result[i].ID,
			Title:     result[i].Title,
			Caption:   result[i].Caption,
			Photo_Url: result[i].Photo_Url,
			User_Id:   result[i].User_Id,
			Create_At: result[i].Create_At,
			User: views.User{
				Email:    email,
				Username: username,
			},
		}
		a[i] = tambah
	}

	WriteJsonResponse_GetPhoto(ctx, &views.Get_Photos{
		Message_Action: "SUCCESS",
		Status:         http.StatusCreated,
		Message_Data:   a,
	})
}

// CreateResource godoc
// @Summary Update Photo
// @Description Update Photo
// @Tags PHOTO
// @Accept  json
// @Produce      json
// @Param Photo body views.Swagger_Comment_Register_Put true "Update Photo"
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Param photoId path int true "Id Photo"
// @Success 200  {object} string "success"
// @Router /photos/{photoId} [put]
func (postgres *HandlersController) Put_Photos(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	println(photoId)
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

	var key_data views.Request_Photos

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
	println("Title: %s", key_data.Title)
	println("Caption: %s", key_data.Caption)
	println("Photo_Url: %s", key_data.Photo_Url)

	// querry user data
	var result models.User
	postgres.db.Table("users").Select("id").Where("email = ?", email).Scan(&result)
	println(result.ID)

	//// Update Data
	postgres.db.Model(&models.Photo{}).Where("id = ?", photoId).Updates(
		models.Photo{
			Title:     key_data.Title,
			Caption:   key_data.Caption,
			Photo_Url: key_data.Photo_Url,
			Update_At: time.Now(),
		})

	intVar, err := strconv.Atoi(photoId)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	WriteJsonResponse_PutPhoto(ctx, &views.Put_Photos{
		Message_Action: "SUCCESS",
		Status:         http.StatusCreated,
		Message_Data: views.Put_Photos_Data{
			ID:        intVar,
			Title:     key_data.Title,
			Caption:   key_data.Caption,
			Photo_Url: key_data.Photo_Url,
			User_Id:   result.ID,
			Update_At: time.Now(),
		},
	})
}

// CreateResource godoc
// @Summary Delete Photo
// @Description Delete Photo
// @Tags PHOTO
// @Accept  json
// @Produce      json
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Param photoId path int true "Id Photo"
// @Success 200  {object} string "success"
// @Router /photos/{photoId} [delete]
func (postgres *HandlersController) Delete_Photos(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	println(photoId)
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
	println(photoId)

	//Query Photo ID
	var result_sm models.Photo
	postgres.db.Table("photos").Where("id = ? AND user_id = ?", photoId, result.ID).Find(&result_sm)
	if result_sm.ID == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Photo ID Not Found on this Account!",
			},
		})
		return
	}

	err := postgres.db.Where("id = ?", photoId).Delete(&models.Photo{}).Error
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Delete Photo Failed",
			},
		})
		return
	}

	WriteJsonResponse_Delete(ctx, &views.Resp_Delete{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data: views.Data_Delete{
			Message: "Your Photo has been successfully deleted!!",
		},
	})
}
