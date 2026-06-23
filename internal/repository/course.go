package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jmoiron/sqlx"
)

type PostgresCourseRepo struct {
	db *sqlx.DB
}

func NewPostgresCourseRepo(dbObj *sqlx.DB) *PostgresCourseRepo {
	repo := &PostgresCourseRepo{
		db: dbObj,
	}

	return repo
}

func (pcr *PostgresCourseRepo) GetAll() ([]model.Course, error) {
	coursesSlice := make([]model.Course, 0)

	query := `
		SELECT id, title, price, is_active, created_at, updated_at
		FROM courses
		ORDER BY created_at
	`

	err := pcr.db.Select(&coursesSlice, query)
	if err != nil {
		return nil, err
	}

	return coursesSlice, nil
}

func (pcr *PostgresCourseRepo) GetByID(id int) (model.Course, error) {
	var course model.Course

	query := `
		SELECT id, title, price, is_active, created_at, updated_at
		FROM courses
		WHERE id = $1
		LIMIT 1
	`

	err := pcr.db.Get(&course, query, id)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (pcr *PostgresCourseRepo) Delete(id int) error {
	query := `
		DELETE FROM courses
		WHERE id = $1
	`

	_, err := pcr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (pcr *PostgresCourseRepo) Create(course model.Course) (int, error) {

	query := `
		INSERT INTO courses (
		    title, price, is_active, created_at, updated_at
		) VALUES (
		    :title, :price, :is_active, :created_at, :updated_at
		)
		RETURNING id
	`

	rows, err := pcr.db.NamedQuery(query, course)
	if err != nil {
		return 0, err
	}

	var id int
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (pcr *PostgresCourseRepo) Update(id int, input model.UpdateCourse) (int, error) {
	var setParts []string
	var args []any
	argID := 1

	if input.Title != nil {
		setParts = append(setParts, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Price != nil {
		setParts = append(setParts, fmt.Sprintf("price=$%d", argID))
		args = append(args, *input.Price)
		argID++
	}

	if input.IsActive != nil {
		setParts = append(setParts, fmt.Sprintf("is_active=$%d", argID))
		args = append(args, *input.IsActive)
		argID++
	}

	if len(setParts) == 0 {
		return 0, errors.New("no fields to update")
	}

	setParts = append(setParts, fmt.Sprintf("updated_at=$%d", argID))
	args = append(args, time.Now())
	argID++

	query := fmt.Sprintf(`
		UPDATE courses
		SET %s
		WHERE id = $%d
	`, strings.Join(setParts, ", "), argID)

	args = append(args, id)

	_, err := pcr.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}
