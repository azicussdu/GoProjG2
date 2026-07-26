package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
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

func (pcr *PostgresCourseRepo) GetAll(ctx context.Context) ([]model.Course, error) {
	coursesSlice := make([]model.Course, 0)

	query := `
		SELECT id, title, description, price, level, is_active, teacher_id, created_at, updated_at
		FROM courses
		WHERE deleted_at IS NULL
		ORDER BY created_at
	`

	err := pcr.db.SelectContext(ctx, &coursesSlice, query)
	if err != nil {
		return nil, apperrors.Internal("failed to get courses", err)
	}

	return coursesSlice, nil
}

func (pcr *PostgresCourseRepo) GetByID(ctx context.Context, id int) (model.Course, error) {
	var course model.Course

	query := `
		SELECT id, title, description, price, level, is_active, teacher_id, created_at, updated_at
		FROM courses
		WHERE id = $1 AND deleted_at IS NULL
		LIMIT 1
	`

	err := pcr.db.GetContext(ctx, &course, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Course{}, apperrors.NotFound("course with ID not found", err)
		}
		return model.Course{}, apperrors.Internal("failed to get course", err)
	}

	return course, nil
}

func (pcr *PostgresCourseRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := `
		UPDATE courses
		SET deleted_at = NOW(),
		    updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return apperrors.Internal("failed to delete course", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return apperrors.Internal("failed to delete course", err)
	}

	if affectedRows == 0 {
		return apperrors.NotFound("course with ID not found", nil)
	}

	return nil
}

func (pcr *PostgresCourseRepo) Create(ctx context.Context, course model.Course) (int, error) {

	query := `
		INSERT INTO courses (
		    title, description, price, level, is_active, teacher_id, created_at, updated_at
		) VALUES (
		    :title, :description, :price, :level, :is_active, :teacher_id, :created_at, :updated_at
		)
		RETURNING id
	`

	rows, err := pcr.db.NamedQueryContext(ctx, query, course)
	if err != nil {
		return 0, apperrors.Internal("failed to created course", err)
	}

	var id int
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, apperrors.Internal("failed to created course", err)
		}
	}

	return id, nil
}

func (pcr *PostgresCourseRepo) Update(ctx context.Context, id int, input model.UpdateCourse) (int, error) {
	var setParts []string
	var args []any
	argID := 1

	if input.Title != nil {
		setParts = append(setParts, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setParts = append(setParts, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	if input.Price != nil {
		setParts = append(setParts, fmt.Sprintf("price=$%d", argID))
		args = append(args, *input.Price)
		argID++
	}

	if input.Level != nil {
		setParts = append(setParts, fmt.Sprintf("level=$%d", argID))
		args = append(args, *input.Level)
		argID++
	}

	if input.IsActive != nil {
		setParts = append(setParts, fmt.Sprintf("is_active=$%d", argID))
		args = append(args, *input.IsActive)
		argID++
	}

	if input.TeacherID != nil {
		setParts = append(setParts, fmt.Sprintf("teacher_id=$%d", argID))
		args = append(args, *input.TeacherID)
		argID++
	}

	if len(setParts) == 0 {
		return 0, apperrors.BadRequest("no fields to update", nil)
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

	result, err := pcr.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, apperrors.Internal("failed to update course", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, apperrors.Internal("failed to update course", err)
	}

	if affectedRows == 0 {
		return 0, apperrors.NotFound("course with ID not found", nil)
	}

	return id, nil
}
