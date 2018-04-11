package model

import (
	"github.com/gin-gonic/gin"
	"github.com/gunkan-s/barnament/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InsertCocktail(db *gorm.DB, cocktail types.Cocktail, base types.Base, timbers []types.Timber) {
	db.Create(&cocktail)
	db.Create(&base)
	for _, timber := range timbers {
		db.Create(&timber)
	}
}
