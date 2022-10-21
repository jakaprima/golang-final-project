package controllers

import (
	"gorm.io/gorm"
)

type HandlersController struct {
	db *gorm.DB
}

func User_DB_Controller(db *gorm.DB) *HandlersController {
	return &HandlersController{
		db: db,
	}
}
