package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/gunkan-s/barnament/model"
)

func InsertCocktailRouting(r *gin.Engine, orm *gorm.DB) {
	r.POST("/insert_cocktail", func(c *gin.Context) {
		// cocktail := c.PostForm("name")
		// base := c.PostForm("base")
		// timber := c.PostForm("timber")

		c.String(200, "pong")
	})
}