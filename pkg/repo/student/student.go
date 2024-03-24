package student

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string `gorm:"column:name"  json:"name"`
	Age   int    `gorm:"column:age"   json:"age"`
	Class int    `gorm:"column:class" json:"class"`
	Email string `gorm:"column:email"  json:"email"`
}

func (s *Student) TableName() string {
	return "student"
}
