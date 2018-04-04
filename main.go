package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gunkan-s/barnament/controller"
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

	// "barnament:barnament@tcp(localhost)/barnament"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	orm := gormConnect()
	r := setupRouter(orm)

	r.Run(":8080")
}
