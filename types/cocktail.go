package types

import (
	"github.com/jinzhu/gorm"
)

type Cocktail struct {
	gorm.Model
	Name        string
	Description string
}

type Base struct {
	gorm.Model
	Cocktail Cocktail
	Name     string
	Vol      string
}

type Timber struct {
	gorm.Model
	Cocktail Cocktail
	Name     string
	Vol      string
}
