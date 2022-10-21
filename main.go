package main

import (
	"finalproject/controllers"
	"finalproject/db"
	"log"
	"net/http"
	"os"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "finalproject/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title My Grams APP
// @description final project golang jaka
// @version v2.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host golang-final-project-production.up.railway.app
// @contact.name Jaka Prima Maulana
// @contact.email jakaprima123@gmail.com

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var PORT = os.Getenv("PORT")
	db := db.ConnectGorm()

	//// Running Service User ////
	serviceController := controllers.User_DB_Controller(db)

	// ------------------------------------------- ROUTER
	user_service := UserRouter(serviceController)
	// ------------------------------------------- START SERVER
	user_service.Start(":" + PORT)
	// user_service.Start(":4000")
}

type Router struct {
	control *controllers.HandlersController
}

func UserRouter(control *controllers.HandlersController) *Router {
	return &Router{control: control}
}

func (r *Router) Start(port string) {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to MyGram"})
	})

	///// User Handlers /////////
	router.POST("/users/register", r.control.Register_User)
	router.POST("/users/login", r.control.Login_User)
	router.PUT("/users", r.control.PUT_User)
	router.DELETE("/users", r.control.Delete_User)

	router.POST("/photos", r.control.Post_Photos)
	router.GET("/photos", r.control.Get_Photos)
	router.PUT("/photos/:photoId", r.control.Put_Photos)
	router.DELETE("/photos/:photoId", r.control.Delete_Photos)

	router.POST("/comments", r.control.Comments_Post)
	router.GET("/comments", r.control.Comment_Get)
	router.PUT("/comments/:commentId", r.control.Comment_Put)
	router.DELETE("/comments/:commentId", r.control.Comment_Delete)

	router.POST("/socialmedias", r.control.SocialMedias_Post)
	router.GET("/socialmedias", r.control.SocialMedias_Get)
	router.PUT("/socialmedias/:socialMediaId", r.control.SocialMedias_Put)
	router.DELETE("/socialmedias/:socialMediaId", r.control.SocialMedias_Delete)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(port)
}
