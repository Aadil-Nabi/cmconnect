package models

import "gorm.io/gorm"

type Identity struct {
	gorm.Model
	IdentityNumber uint
	Department     string
}
