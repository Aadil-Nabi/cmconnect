package models

import "gorm.io/gorm"

type Identity struct {
	gorm.Model
	IdentityNumber string `gorm:"identitynumber"`
	Department     string `gorm:"department"`
}
