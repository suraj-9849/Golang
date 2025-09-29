package storage

import "github.com/suraj-9849/Golang.git/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetList() ([]types.Student, error)
	DeleteTheStudentByID(id int64) (types.Student, error)
	UpdateById(id int64, name string, email string, age int) (types.Student, error)
}
