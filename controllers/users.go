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
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(userid int) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd34tg2y4j7j") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////// HANDLERS FOR USER ////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////

// CreateResource godoc
// @Summary Creates User account
// @Description Register User
// @Tags USER
// @Accept  json
// @Produce      json
// @Param user body views.Swagger_User_Register_Post true "Register User"
// @Success 201  {object} string "success"
// @Router /users/register [post]
func (postgres *HandlersController) Register_User(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	var key_data views.Request_Register

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "failed to registrate user!!",
			},
		})
		return
	}

	println(key_data.Username)
	println(key_data.Email)
	println(key_data.Password)
	println(key_data.Age)

	valid(key_data.Email)

	// Email type Validation
	if valid(key_data.Email) != true {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "email format not valid!!",
			},
		})
		return
	}

	// Check Email registered Database
	email_check := postgres.db.Where("email = ?", key_data.Email).First(&models.User{}).Error
	fmt.Println(email_check)
	if email_check == nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "email already registered for another account!",
			},
		})
		return
	}

	// Username Validation
	if key_data.Username == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "username field is empty!",
			},
		})
		return
	}
	username_check := postgres.db.Where("username = ?", key_data.Username).First(&models.User{}).Error
	fmt.Println(email_check)
	if username_check == nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "username alreade use for another acoount!",
			},
		})
		return
	}

	// Password Validation :
	if key_data.Password == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "password field can't be empty!",
			},
		})
		return
	}
	password_length := len(key_data.Password)
	if password_length < 6 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "password lenght must more than 6 characters!!",
			},
		})
		return
	}
	//encript Password with bycript
	hash, _ := HashPassword(key_data.Password)

	// validation Age
	age := key_data.Age
	if age < 9 {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "age must be more than 9 years old!",
			},
		})
		return
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	err_create := postgres.db.Create(&models.User{
		ID:        r1.Int() / 100000,
		Username:  key_data.Username,
		Email:     key_data.Email,
		Password:  hash,
		Age:       key_data.Age,
		Create_At: time.Now(),
		Update_At: time.Now(),
	}).Error
	if err_create != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "register user faile to store db!",
			},
		})
		return
	}

	WriteJsonResponse_Succes(ctx, &views.Resp_Register_Success{
		Message_Action: "SUCCESS",
		Status:         http.StatusCreated,
		Message_Data: views.Data_Register{
			Age:       key_data.Age,
			Email:     key_data.Email,
			ID:        r1.Int() / 100000,
			Username:  key_data.Username,
			Create_At: time.Now(),
			Update_At: time.Now(),
		},
	})
}

// CreateResource godoc
// @Summary Login Account
// @Description Login Account
// @Tags USER
// @Accept  json
// @Produce      json
// @Param user body views.Swagger_User_Login_Post true "Login User"
// @Success 200  {object} string "success"
// @Router /users/login [post]
func (postgres *HandlersController) Login_User(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	var key_data views.Request_Login

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "login user failed!",
			},
		})
		return
	}

	println(key_data.Email)
	println(key_data.Password)

	// Email type Validation
	if valid(key_data.Email) != true {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "email format not valid!!",
			},
		})
		return
	}

	// Password and Email Login Check Acount
	var s sql.NullString

	postgres.db.Select("password").Where("email = ?", key_data.Email).First(&models.User{}).Scan(&s)
	password_from_db := s.String
	fmt.Printf("%s \n", password_from_db)

	match := CheckPasswordHash(key_data.Password, password_from_db)
	fmt.Println("Match:   ", match)
	if match != true {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "email and password does not match!",
			},
		})
		return
	}

	// generate JWT-TOKEN
	var result models.User
	postgres.db.Table("users").Select("email", "username").Where("email = ?", key_data.Email).Scan(&result)
	println(result.Age)

	token, err := GenerateJWT(result.Email, result.Username)
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "failed generate Token!",
			},
		})
		return
	}

	WriteJsonResponse_Login(ctx, &views.Resp_Login{
		Message_Action: "SUCCESS",
		Status:         http.StatusOK,
		Message_Data: views.Token{
			Token: token,
		},
	})
}

// CreateResource godoc
// @Summary Update Data Account
// @Description Update Data Account
// @Tags USER
// @Accept  json
// @Produce      json
// @Param User body views.Swagger_User_Update_Put true "Update Data User"
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Param userid query int  true  "User ID"
// @Success 200  {object} string "success"
// @Router /users [put]
func (postgres *HandlersController) PUT_User(ctx *gin.Context) {
	// Check Authorization
	tokenString := ctx.GetHeader("Authorization")
	fmt.Println(tokenString)
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
	jwtString := strings.ReplaceAll(tokenString, "Bearer ", "")
	err1 := ValidateToken(tokenString)
	if err1 != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
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

	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	var key_data views.Request_Update

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "update user failed!",
			},
		})
		return
	}

	// Email type Validation
	if valid(key_data.Email) != true {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "Email format not valid!",
			},
		})
		return
	}

	// Check Email registered Database
	email_check := postgres.db.Where("email = ?", key_data.Email).First(&models.User{}).Error
	fmt.Println(email_check)
	if email_check == nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "Email Already Registered!",
			},
		})
		return
	}

	// Username Validation
	if key_data.Username == "" {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "Username Field is empty!",
			},
		})
		return
	}
	username_check := postgres.db.Where("username = ?", key_data.Username).First(&models.User{}).Error
	fmt.Println(email_check)
	if username_check == nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "VALIDATION_ERROR",
			Status:         http.StatusUnprocessableEntity,
			Message_Data: views.Message{
				Message: "Username Already used for another account!",
			},
		})
		return
	}

	// querry user data
	var result models.User
	postgres.db.Table("users").Select("age", "email", "id", "username").Where("email = ?", email).Scan(&result)
	println(result.ID)

	//// Update Data
	postgres.db.Model(&models.User{}).Where("id = ?", result.ID).Updates(
		models.User{
			Username:  key_data.Username,
			Email:     key_data.Email,
			Update_At: time.Now(),
		})

	WriteJsonResponse_Put(ctx, &views.Resp_Put{
		Message_Action: "SUCCESS",
		Status:         http.StatusCreated,
		Message_Data: views.Put{
			Age:       result.Age,
			Email:     result.Email,
			ID:        result.ID,
			Username:  result.Username,
			Update_At: time.Now(),
		},
	})
}

// CreateResource godoc
// @Summary Delete Data Account
// @Description Delete Data Account
// @Tags USER
// @Accept  json
// @Produce      json
// @Param Authorization header string  true  "Token Barier example: 'Bearer 12355f32r'"
// @Success 200  {object} string "success"
// @Router /users [delete]
func (postgres *HandlersController) Delete_User(ctx *gin.Context) {
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

	err_social_media := postgres.db.Where("user_id = ?", result.ID).Delete(&models.SocialMedia{}).Error
	if err_social_media != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Delete User Failed",
			},
		})
		return
	}

	err_comment := postgres.db.Where("user_id = ?", result.ID).Delete(&models.Comment{}).Error
	if err_comment != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Delete Comment Failed",
			},
		})
		return
	}

	err_photo := postgres.db.Where("user_id = ?", result.ID).Delete(&models.Photo{}).Error
	if err_photo != nil {
		WriteJsonResponse_Failed(ctx, &views.Failed{
			Message_Action: "GENERAL_REQUEST_ERROR",
			Status:         http.StatusInternalServerError,
			Message_Data: views.Message{
				Message: "Delete Photo Failed",
			},
		})
		return
	}

	err := postgres.db.Where("username = ?", username).Delete(&models.User{}).Error
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
			Message: "Your account has been successfully deleted",
		},
	})
}
