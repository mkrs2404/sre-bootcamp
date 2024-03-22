package student

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}
