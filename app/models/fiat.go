package models

import "gorm.io/gorm"

type Fiat struct {
	gorm.Model
	Name   string
	Sign   string
	Symbol string
}
