package student

import uuid "github.com/satori/go.uuid"

func ToStudent(studentDTO StudentDTO) Student {
	return Student{Identifier: studentDTO.Identifier, Name: studentDTO.Name, Email: studentDTO.Email, UKT: studentDTO.UKT}
}

func ToStudent2(studentDTO CreateStudentDTO) Student {
	identifier := uuid.NewV4().String()
	return Student{Identifier: identifier, Name: studentDTO.Name, Email: studentDTO.Email, UKT: studentDTO.UKT}
}

func ToStudentDTO(student Student) StudentDTO {
	return StudentDTO{ID: student.ID, Identifier: student.Identifier, Name: student.Name, Email: student.Email, UKT: student.UKT}
}

func ToStudentDTOs(students []Student) []StudentDTO {
	studentdtos := make([]StudentDTO, len(students))

	for i, itm := range students {
		studentdtos[i] = ToStudentDTO(itm)
	}

	return studentdtos
}