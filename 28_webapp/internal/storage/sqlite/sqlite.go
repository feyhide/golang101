package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/feyhide/golang101/28_webapp/internal/config"
	"github.com/feyhide/golang101/28_webapp/internal/types"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	//go doesnt have constructor following func will help us
	//simulating that and returns us a instance
	db, err := sql.Open("sqlite", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	age INTEGER,
	email TEXT
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

// this will fix implementation error in main.go sending storage in new()
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	statement, err := s.Db.Prepare("INSERT INTO students (name,email,age) VALUES (?,?,?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare: %w", err)

	}
	defer statement.Close()

	result, err := statement.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetById(id int64) (types.Student, error) {
	statement, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	var student types.Student

	err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Age, &student.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student not found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetList() ([]types.Student, error) {
	statement, err := s.Db.Prepare("SELECT id, name, age, email  FROM students")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (s *Sqlite) UpdateStudent(id int64, name string, age int, email string) (types.Student, error) {
	statement, err := s.Db.Prepare("UPDATE students SET name = ?, age = ?, email = ? WHERE id = ?")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare update statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(name, age, email, id)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to execute update statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to fetch rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return types.Student{}, fmt.Errorf("no student found with id %d", id)
	}

	var updatedStudent types.Student
	err = s.Db.QueryRow("SELECT id, name, age, email FROM students WHERE id = ?", id).Scan(&updatedStudent.Id, &updatedStudent.Name, &updatedStudent.Age, &updatedStudent.Email)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to fetch updated student: %w", err)
	}

	return updatedStudent, nil
}

func (s *Sqlite) DeleteStudent(id int64) (types.Student, error) {
	statement, err := s.Db.Prepare("SELECT id, name, age, email FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare select statement: %w", err)
	}
	defer statement.Close()

	var student types.Student
	err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Age, &student.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student not found with id %d", id)
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	deleteStatement, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare delete statement: %w", err)
	}
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec(id)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to execute delete statement: %w", err)
	}

	return student, nil
}
