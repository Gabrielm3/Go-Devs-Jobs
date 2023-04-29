package handler

import (
	"gorm.io/gorm"

	"github.com/Gabrielm3/Go-Devs-Jobs/config"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
