package student

type StudentService struct {
	StudentRepository StudentRepository
}

func ProvideStudentService(p StudentRepository) StudentService {
	return StudentService{StudentRepository: p}
}

func (p *StudentService) FindAll() []Student {
	return p.StudentRepository.FindAll()
}

func (p *StudentService) FindByID(id uint) Student {
	return p.StudentRepository.FindById(id)
}

func (p *StudentService) Save(student Student) Student {
	p.StudentRepository.Save(student)
	return student
}

func (p *StudentService) Delete(student Student) {
	p.StudentRepository.Delete(student)
}