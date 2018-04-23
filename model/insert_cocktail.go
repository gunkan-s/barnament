package model

import (
	"github.com/gunkan-s/barnament/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Cocktail struct{}

func CreateCocktail(db *gorm.DB, cocktail types.Cocktail, base types.Base, timbers []types.Timber) error {
	db.Create(&cocktail)
	db.Create(&base)
	for _, timber := range timbers {
		db.Create(&timber)
	}
	return nil
}
