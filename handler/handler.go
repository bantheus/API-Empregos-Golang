package handler

import (
	"github.com/bantheus/API-Empregos-Golang/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostgreSQL()
}
