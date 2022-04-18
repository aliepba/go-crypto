package models

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Coin         string
	Symbol       string
	Description  string
	Logo         string
	UserID       int
	CategoryId   int
	MetadataCoin MetadataCoin
	User         User
	Category     Category
}
