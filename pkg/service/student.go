package service

import "github.com/mkrs2404/sre-bootcamp/pkg/repo/student"

type StudentService interface {
	Create(a *student.Student) error
	GetAll() ([]*student.Student, error)
	GetByID(id int) (*student.Student, error)
	Update(a *student.Student) error
	Delete(a *student.Student) error
}

type studentService struct {
	studentRepository student.Repository
}

func NewStudentService(ar student.Repository) StudentService {
	return &studentService{ar}
}

func (s studentService) GetAll() ([]*student.Student, error) {
	return s.studentRepository.GetAll()
}

func (s studentService) GetByID(id int) (*student.Student, error) {
	return s.studentRepository.GetByID(id)
}

func (s studentService) Create(a *student.Student) error {
	return s.studentRepository.Create(a)
}

func (s studentService) Update(a *student.Student) error {
	return s.studentRepository.Update(a)
}

func (s studentService) Delete(a *student.Student) error {
	return s.studentRepository.Delete(a)
}