package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gunkan-s/barnament/controller"
	"github.com/gunkan-s/barnament/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupRouter(orm *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("view/html/*.html")
	r.GET("/insert_cocktail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "insert_cocktail.html", gin.H{})
	})
	controller.InsertCocktailRouting(r, orm)

	return r
}

func gormConnect() *gorm.DB {
	u, _ := url.Parse(os.Getenv("CLEARDB_DATABASE_URL"))
	CONNECT := fmt.Sprintf("%s@tcp(%s:3306)%s", u.User.String(), u.Host, u.Path)
	db, err := gorm.Open("mysql", CONNECT)
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

	r.Run(":" + os.Getenv("PORT"))
}
