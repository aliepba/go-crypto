package models

import "gorm.io/gorm"

type MetadataCoin struct {
	gorm.Model
	Website      string
	TechnicalDoc string
	Twitter      string
	Reddit       string
	MessageBoard string
	SourceCode   string
	CoinID       int
	UserID       int
	User         User
}
