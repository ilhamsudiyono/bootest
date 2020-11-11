package dao

import (
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
)

var g *gorm.DB

func SetDao(gDB *gorm.DB) {
	g = gDB
}

func catchError(e *error) {
	config.CatchError(e)
}
