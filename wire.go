package main

import (
	"Go/student"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func _(db *gorm.DB) student.StudentAPI {
	wire.Build(student.ProvideStudentRepository, student.ProvideStudentService, student.ProvideStudentAPI)
	return student.StudentAPI{}
}