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

type PostgresLessonRepo struct {
	db *sqlx.DB
}

func NewPostgresLessonRepo(db *sqlx.DB) *PostgresLessonRepo {
	return &PostgresLessonRepo{db: db}
}

func (r *PostgresLessonRepo) GetByCourseID(ctx context.Context, courseID int) ([]model.Lesson, error) {
	lessons := make([]model.Lesson, 0)

	query := `
		SELECT id, course_id, title, content, position, created_at, updated_at
		FROM lessons
		WHERE course_id = $1 AND deleted_at IS NULL
		ORDER BY position
	`

	err := r.db.SelectContext(ctx, &lessons, query, courseID)
	if err != nil {
		return nil, apperrors.Internal("failed to get lessons", err)
	}

	return lessons, nil
}

func (r *PostgresLessonRepo) GetByID(ctx context.Context, id int) (model.Lesson, error) {
	var lesson model.Lesson

	query := `
		SELECT id, course_id, title, content, position, created_at, updated_at
		FROM lessons
		WHERE id = $1 AND deleted_at IS NULL
		LIMIT 1
	`

	err := r.db.GetContext(ctx, &lesson, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Lesson{}, apperrors.NotFound("lesson with ID not found", err)
		}
		return model.Lesson{}, apperrors.Internal("failed to get lesson", err)
	}

	return lesson, nil
}

func (r *PostgresLessonRepo) Create(ctx context.Context, lesson model.Lesson) (int, error) {
	query := `
		INSERT INTO lessons (course_id, title, content, position, created_at, updated_at)
		VALUES (:course_id, :title, :content, :position, :created_at, :updated_at)
		RETURNING id
	`

	rows, err := r.db.NamedQueryContext(ctx, query, lesson)
	if err != nil {
		return 0, apperrors.Internal("failed to create lesson", err)
	}

	var id int
	if rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return 0, apperrors.Internal("failed to create lesson", err)
		}
	}

	return id, nil
}

func (r *PostgresLessonRepo) Update(ctx context.Context, id int, input model.UpdateLesson) (int, error) {
	var setParts []string
	var args []any
	argID := 1

	if input.Title != nil {
		setParts = append(setParts, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Content != nil {
		setParts = append(setParts, fmt.Sprintf("content=$%d", argID))
		args = append(args, *input.Content)
		argID++
	}

	if input.Position != nil {
		setParts = append(setParts, fmt.Sprintf("position=$%d", argID))
		args = append(args, *input.Position)
		argID++
	}

	if len(setParts) == 0 {
		return 0, apperrors.BadRequest("no fields to update", nil)
	}

	setParts = append(setParts, fmt.Sprintf("updated_at=$%d", argID))
	args = append(args, time.Now())
	argID++

	query := fmt.Sprintf(`
		UPDATE lessons
		SET %s
		WHERE id = $%d AND deleted_at IS NULL
	`, strings.Join(setParts, ", "), argID)

	args = append(args, id)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, apperrors.Internal("failed to update lesson", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, apperrors.Internal("failed to update lesson", err)
	}

	if affectedRows == 0 {
		return 0, apperrors.NotFound("lesson with ID not found", nil)
	}

	return id, nil
}

func (r *PostgresLessonRepo) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE lessons
		SET deleted_at = NOW(),
		    updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return apperrors.Internal("failed to delete lesson", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return apperrors.Internal("failed to delete lesson", err)
	}

	if affectedRows == 0 {
		return apperrors.NotFound("lesson with ID not found", nil)
	}

	return nil
}
