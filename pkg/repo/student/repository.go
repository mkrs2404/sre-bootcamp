package student

// Repository is an interface that defines the methods for querying
// the database for students.
type Repository interface {
	GetAll() ([]Student, error)
	GetByID(int) (*Student, error)
	Create(*Student) error
	Update(*Student) error
	Delete(int) error
}
