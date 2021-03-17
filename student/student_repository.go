package student

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type StudentRepository struct {
	DB *gorm.DB
}

func ProvideStudentRepository(DB *gorm.DB) StudentRepository {
	return StudentRepository{DB: DB}
}

func (p *StudentRepository) FindAll() []Student {
	var students []Student
	p.DB.Find(&students)
	return students
}

func (p *StudentRepository) FindById(id uint) Student {
	var student Student
	p.DB.First(&student, id)
	return student
}

func (p *StudentRepository) Save(student Student) Student {
	p.DB.Save(&student)
	return student
}

func (p *StudentRepository) Delete(student Student) {
	p.DB.Delete(&student)
}