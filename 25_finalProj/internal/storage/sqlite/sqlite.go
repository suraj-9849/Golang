package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/suraj-9849/Golang.git/internal/config"
	"github.com/suraj-9849/Golang.git/internal/types"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO STUDENTS (name, email, age) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM STUDENTS WHERE id=? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()
	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %d", id)
		}
		return types.Student{}, err
	}
	return student, nil
}
func (s *Sqlite) GetList() ([]types.Student, error) {
	rows, err := s.Db.Query("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (s *Sqlite) DeleteTheStudentByID(id int64) (types.Student, error) {
	// First, fetch the student
	var student types.Student
	err := s.Db.QueryRow("SELECT id, name, email, age FROM students WHERE id = ?", id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %d", id)
		}
		return types.Student{}, err
	}

	// Now, delete the student
	_, err = s.Db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return types.Student{}, err
	}
	return student, nil
}

func (s *Sqlite) UpdateById(id int64, name string, email string, age int) (types.Student, error) {
	// First, check if student exists
	var student types.Student
	err := s.Db.QueryRow("SELECT id, name, email, age FROM students WHERE id = ?", id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %d", id)
		}
		return types.Student{}, err
	}

	// Now, update the student
	_, err = s.Db.Exec("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?", name, email, age, id)
	if err != nil {
		return types.Student{}, err
	}

	// Return the updated student
	student.Name = name
	student.Email = email
	student.Age = age
	return student, nil
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{Db: db}, nil
}
