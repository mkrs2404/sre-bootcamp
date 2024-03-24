package service

import (
	"github.com/mkrs2404/sre-bootcamp/pkg/repo/student"
)

type StudentService interface {
	Create(s *student.Student) error
	GetAll() ([]student.Student, error)
	GetByID(id int) (*student.Student, error)
	Update(s *student.Student) error
	Delete(id int) error
}

type studentService struct {
	studentRepository student.Repository
}

func NewStudentService(sr student.Repository) StudentService {
	return &studentService{sr}
}

func (s studentService) GetAll() ([]student.Student, error) {
	return s.studentRepository.GetAll()
}

func (s studentService) GetByID(id int) (*student.Student, error) {
	return s.studentRepository.GetByID(id)
}

func (s studentService) Create(st *student.Student) error {
	return s.studentRepository.Create(st)
}

func (s studentService) Update(st *student.Student) error {
	_, err := s.studentRepository.GetByID(int(st.ID))
	if err != nil {
		return err
	}
	return s.studentRepository.Update(st)
}

func (s studentService) Delete(id int) error {
	_, err := s.studentRepository.GetByID(id)
	if err != nil {
		return err
	}
	return s.studentRepository.Delete(id)
}
