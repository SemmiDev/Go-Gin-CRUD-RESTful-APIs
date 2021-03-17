package student

import (
	"strconv"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type StudentAPI struct {
	StudentService StudentService
}

func ProvideStudentAPI(p StudentService) StudentAPI {
	return StudentAPI{StudentService: p}
}

func (p *StudentAPI) FindAll(c *gin.Context) {
	students := p.StudentService.FindAll()

	c.JSON(http.StatusOK, gin.H{"students": ToStudentDTOs(students)})
}

func (p *StudentAPI) FindByID(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	student := p.StudentService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"student": ToStudentDTO(student)})
}

func (p *StudentAPI) Create(c *gin.Context) {
	var createStudentDTO CreateStudentDTO
	err := c.BindJSON(&createStudentDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdStudent := p.StudentService.Save(ToStudent2(createStudentDTO))
	c.JSON(http.StatusOK, gin.H{"student": ToStudentDTO(createdStudent)})
}

func (p *StudentAPI) Update(c *gin.Context) {
	var studentDTO UpdateStudentDTO
	err := c.BindJSON(&studentDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ :=  strconv.Atoi(c.Param("id"))
	student := p.StudentService.FindByID(uint(id))
	if student == (Student{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	student.Name = studentDTO.Name
	student.Email = studentDTO.Email
	p.StudentService.Save(student)

	c.Status(http.StatusOK)
}

func (p *StudentAPI) Delete(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	student := p.StudentService.FindByID(uint(id))
	if student == (Student{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.StudentService.Delete(student)
	c.Status(http.StatusOK)
}