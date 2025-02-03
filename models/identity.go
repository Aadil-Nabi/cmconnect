package models

import "gorm.io/gorm"

type Identity struct {
	gorm.Model
	IdentityNumber string `gorm:"identitynumber"`
	Mac            string `gorm:"mac"`
	Department     string `gorm:"department"`
}
