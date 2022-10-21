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
// @Summary Post Comments
// @Description Post Comments
// @Tags COMMENTS
// @Accept  json
// @Produce      json
// @Param Comments body views.Swagger_Comment_Register_Post true "Post Comments"
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Success 200  {object} string "success"
// @Router /comments [post]
func (postgres *HandlersController) Comments_Post(ctx *gin.Context) {
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

	var key_data views.Request_Coments

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
	println("Message: ", key_data.Message)
	println("Photo_Id: ", key_data.Photo_Id)

	// Tittle and Photo_Url Validation
	if key_data.Message == "" || key_data.Photo_Id == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "Message or Photo_Id field is empty!",
			},
		})
		return
	}

	// Check data from table photo
	var poto models.Photo
	postgres.db.Table("photos").Select("id").Where("id = ?", key_data.Photo_Id).Scan(&poto)
	println(poto.ID)
	if poto.ID == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "Photo ID Not Found!",
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

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// Create/Insert Photo to databases
	err_photo := postgres.db.Create(&models.Comment{
		ID:        r1.Int() / 100000,
		User_Id:   result.ID,
		Photo_Id:  key_data.Photo_Id,
		Message:   key_data.Message,
		Create_At: time.Time{},
		Update_At: time.Time{},
	}).Error

	if err_photo != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "failed to post Comment to db",
			},
		})
		return
	}

	WriteJsonResponse_PostComments(ctx, &views.Comments_Post{
		Message_Action: "SUCCESS",
		Status:         http.StatusCreated,
		Message_Data: views.Comments_Post_Data{
			ID:        r1.Int() / 100000,
			Message:   key_data.Message,
			Photo_Id:  key_data.Photo_Id,
			User_Id:   result.ID,
			Create_At: time.Time{},
		},
	})
}

// CreateResource godoc
// @Summary Get Comments
// @Description Get Comments
// @Tags COMMENTS
// @Accept  json
// @Produce      json
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Success 200  {object} string "success"
// @Router /comments [get]
func (postgres *HandlersController) Comment_Get(ctx *gin.Context) {
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

	var result_comment []models.Comment
	err := postgres.db.Table("comments").Where("user_id = ?", user_id).First(&result_comment).Error
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "No Comment Found!!",
			},
		})
		return
	}
	hasil_comment := postgres.db.Table("comments").Where("user_id = ?", user_id).Find(&result_comment)

	iduser, err := strconv.Atoi(user_id)
	fmt.Println(iduser, err, reflect.TypeOf(iduser))

	var a = make([]interface{}, hasil_comment.RowsAffected)
	for i := 0; i < int(hasil_comment.RowsAffected); i++ {

		var result_foto models.Photo
		postgres.db.Table("photos").Where("id = ?", result_comment[i].Photo_Id).Find(&result_foto)

		tambah := &views.Comments_Put_Data{
			ID:        result_comment[i].ID,
			Message:   result_comment[i].Message,
			Photo_Id:  result_comment[i].Photo_Id,
			User_Id:   result_comment[i].User_Id,
			Create_At: result_comment[i].Create_At,
			Update_at: result_comment[i].Update_At,
			User: views.Coment_User{
				ID:       iduser,
				Email:    email,
				Username: username,
			},
			Photo: views.Coment_Photo{
				ID:        result_foto.ID,
				Title:     result_foto.Title,
				Caption:   result_foto.Caption,
				Photo_Url: result_foto.Photo_Url,
				User_Id:   result_foto.User_Id,
			},
		}
		a[i] = tambah
	}

	WriteJsonResponse_GetComment(ctx, &views.Get_Comment{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data:   a,
	})

	fmt.Println(result_comment[1])
	fmt.Println(hasil_comment)
}

// CreateResource godoc
// @Summary Update Comments
// @Description Update Comments
// @Tags COMMENTS
// @Accept  json
// @Produce      json
// @Param Comments body views.Swagger_Comment_Register_Put true "Update Comments"
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Param commentId path int true "ID Comment"
// @Success 200  {object} string "success"
// @Router /comments/{commentId} [put]
func (postgres *HandlersController) Comment_Put(ctx *gin.Context) {
	CommentId := ctx.Param("commentId")
	println(CommentId)
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

	var key_data views.Request_Coments

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
	println("Message: ", key_data.Message)

	intVar, err := strconv.Atoi(CommentId)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	postgres.db.Model(&models.Comment{}).Where("id = ?", intVar).Updates(
		models.Comment{
			Message:   key_data.Message,
			Update_At: time.Now(),
		})

	var comment_foto models.Comment
	postgres.db.Table("comments").Where("id = ?", intVar).Find(&comment_foto)
	if comment_foto.Photo_Id == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Comment ID Not Found!!",
			},
		})
		return
	}

	var photo_Detail models.Photo
	postgres.db.Table("photos").Where("id = ?", comment_foto.Photo_Id).Find(&photo_Detail)

	WriteJsonResponse_PutComment(ctx, &views.Put_Comment{
		Message_Action: "SUCCESS",
		Status:         http.StatusCreated,
		Message_Data: views.Put_Comment_Data{
			ID:        intVar,
			Title:     photo_Detail.Title,
			Caption:   photo_Detail.Caption,
			Photo_Url: photo_Detail.Photo_Url,
			User_Id:   photo_Detail.User_Id,
			Update_At: time.Now(),
		},
	})
}

// CreateResource godoc
// @Summary Delete Comments
// @Description Delete Comments
// @Tags COMMENTS
// @Accept  json
// @Produce      json
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Param commentId path int true "ID Comment"
// @Success 200  {object} string "success"
// @Router /comments/{commentId} [delete]
func (postgres *HandlersController) Comment_Delete(ctx *gin.Context) {
	CommentId := ctx.Param("commentId")
	println(CommentId)
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
	postgres.db.Table("comments").Where("id = ? AND user_id = ?", CommentId, result.ID).Find(&result_sm)
	if result_sm.ID == 0 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Coment ID Not Found on this Account!",
			},
		})
		return
	}

	err := postgres.db.Where("id = ?", CommentId).Delete(&models.Comment{}).Error
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Delete Comment Failed",
			},
		})
		return
	}

	WriteJsonResponse_Delete(ctx, &views.Resp_Delete{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data: views.Data_Delete{
			Message: "Your comment has been successfully deleted!!",
		},
	})
}
