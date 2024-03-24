package gormimpl

import (
	"github.com/mkrs2404/sre-bootcamp/pkg/db"
	"github.com/mkrs2404/sre-bootcamp/pkg/errors"
	"github.com/mkrs2404/sre-bootcamp/pkg/repo/student"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *db.DB) *studentRepository {
	return &studentRepository{
		db: db.DB,
	}
}

// GetAll returns a list of all students in the database.
func (r *studentRepository) GetAll() ([]student.Student, error) {
	var students []student.Student
	if err := r.db.Find(&students).Error; err != nil {
		logrus.Errorf("unable to fetch the students: %v", err)
		return nil, err
	}

	return students, nil
}

// GetByID returns a student by its ID.
func (r *studentRepository) GetByID(id int) (*student.Student, error) {
	var student student.Student
	if err := r.db.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logrus.Errorf("student %d not found", id)
			return nil, errors.ErrNotFound
		}
		logrus.Errorf("unable to fetch the student %d: %v", id, err)
		return nil, err
	}

	return &student, nil
}

// Create a student in the database.
func (r *studentRepository) Create(s *student.Student) error {
	err := r.db.Save(&s).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			logrus.Errorf("email %s already exists : %v", s.Email, err)
			return errors.ErrEmailAlreadyExists
		}
		logrus.Errorf("unable to create the student: %v", err)
	}
	return err
}

// Update the student in the database.
func (r *studentRepository) Update(s *student.Student) error {
	err := r.db.Updates(&s).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			logrus.Errorf("email %s already exists : %v", s.Email, err)
			return errors.ErrEmailAlreadyExists
		}
		logrus.Errorf("unable to update the student %d: %v", s.ID, err)
	}
	return err
}

// Delete the student from the database.
func (r *studentRepository) Delete(id int) error {
	err := r.db.Unscoped().Delete(&student.Student{}, id).Error
	if err != nil {
		logrus.Errorf("unable to delete the student %d: %v", id, err)
	}
	return err
}
