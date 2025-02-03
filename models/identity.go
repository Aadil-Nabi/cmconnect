package models

import "gorm.io/gorm"

type Identity struct {
	gorm.Model
	IdentityNumber string `gorm:"identitynumber"`
	SecurityPin    string `gorm:"securitypin"`
	Department     string `gorm:"department"`
}
