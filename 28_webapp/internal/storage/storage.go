package storage

import "github.com/feyhide/golang101/28_webapp/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetById(id int64) (types.Student, error)
	GetList() ([]types.Student, error)
	UpdateStudent(id int64, name string, age int, email string) (types.Student, error)
	DeleteStudent(id int64) (types.Student, error)
}
