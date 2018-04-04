package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gunkan-s/barnament/controller"
	"github.com/gunkan-s/barnament/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupRouter(orm *gorm.DB) *gin.Engine {
	r := gin.Default()

	controller.InsertCocktailRouting(r, orm)

	return r
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "bar"
	PASS := "bar"
	PROTOCOL := "tcp(localhost)"
	DBNAME := "barnament"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.AutoMigrate(&types.Cocktail{})
	db.AutoMigrate(&types.Base{})
	db.AutoMigrate(&types.Timber{})

	return db
}

func main() {
	db := gormConnect()
	r := setupRouter(db)

	r.Run(":80")
}
