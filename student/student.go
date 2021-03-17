package student

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Identifier 	string
	Name 		string
	Email 		string
	UKT 		uint
}