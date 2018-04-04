package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gunkan-s/barnament/model"
	"github.com/gunkan-s/barnament/types"
	"github.com/jinzhu/gorm"
	"strconv"
)

func InsertCocktailRouting(r *gin.Engine, db *gorm.DB) {
	r.POST("/insert_cocktail", func(c *gin.Context) {
		cocktail := types.Cocktail{
			Name:        c.PostForm("cocktail-name"),
			Description: c.PostForm("description"),
		}
		base := types.Base{
			Cocktail: cocktail,
			Name:     c.PostForm("base-name"),
			Vol:      c.PostForm("base-vol"),
		}
		timbers := []types.Timber{}
		for i := 1; i <= 5; i++ {
			index := strconv.Itoa(i)
			timbers = append(timbers, types.Timber{
				Cocktail: cocktail,
				Name:     c.PostForm("timber-name-" + index),
				Vol:      c.PostForm("timber-vol-" + index),
			})
		}

		model.InsertCocktail(db, cocktail, base, timbers)

		c.String(200, "pong")
	})
}
