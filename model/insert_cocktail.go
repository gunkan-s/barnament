package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
func validateName(name) {
}

func validateBase(base) {
}

func validateTimber(timber) {
}
*/

func InsertCocktail(gorm *gorm.DB, name string, base string, timber string) {
}