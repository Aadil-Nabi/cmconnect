package models

import "gorm.io/gorm"

type Identity struct {
	gorm.Model
	EmployeeName string `gorm:"employeename"`
	SecurityPin  string `gorm:"securitypin"`
	Department   string `gorm:"department"`
}
