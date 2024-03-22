package gormimpl

import (
	"github.com/mkrs2404/sre-bootcamp/pkg/db"
	"github.com/mkrs2404/sre-bootcamp/pkg/repo/student"
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
func (r *studentRepository) GetAll() ([]*student.Student, error) {
	var animals []*student.Student
	if err := r.db.Find(&animals).Error; err != nil {
		return nil, err
	}

	return animals, nil
}

// GetByID returns a student by its ID.
func (r *studentRepository) GetByID(id int) (*student.Student, error) {
	var student *student.Student
	if err := r.db.First(student, id).Error; err != nil {
		return nil, err
	}

	return student, nil
}

// Create a student in the database.
func (r *studentRepository) Create(a *student.Student) error {
	err := r.db.Save(&a).Error
	return err
}

// Update the student in the database.
func (r *studentRepository) Update(a *student.Student) error {
	err := r.db.Updates(&a).Error
	return err
}

// Delete the student from the database.
func (r *studentRepository) Delete(a *student.Student) error {
	err := r.db.Delete(&a).Error
	return err
}
