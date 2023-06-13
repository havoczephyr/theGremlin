package models

import "gorm.io/gorm"

type Questions struct {
	gorm.Model
	Body   string
	Author string
}
