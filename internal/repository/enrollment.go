package repository

import (
	"context"
	"errors"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

type PostgresEnrollmentRepo struct {
	db *sqlx.DB
}

func NewPostgresEnrollmentRepo(db *sqlx.DB) *PostgresEnrollmentRepo {
	return &PostgresEnrollmentRepo{db: db}
}

func (r *PostgresEnrollmentRepo) Create(ctx context.Context, enrollment model.Enrollment) (int, error) {
	query := `
		INSERT INTO enrollments (student_id, course_id, created_at, updated_at)
		VALUES (:student_id, :course_id, :created_at, :updated_at)
		RETURNING id
	`

	rows, err := r.db.NamedQueryContext(ctx, query, enrollment)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
			return 0, apperrors.Conflict("already enrolled in this course", err)
		}
		return 0, apperrors.Internal("failed to create enrollment", err)
	}

	var id int
	if rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return 0, apperrors.Internal("failed to create enrollment", err)
		}
	}

	return id, nil
}

func (r *PostgresEnrollmentRepo) Delete(ctx context.Context, studentID, courseID int) error {
	query := `
		UPDATE enrollments
		SET deleted_at = NOW(),
		    updated_at = NOW()
		WHERE student_id = $1 AND course_id = $2 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, studentID, courseID)
	if err != nil {
		return apperrors.Internal("failed to leave course", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return apperrors.Internal("failed to leave course", err)
	}

	if affectedRows == 0 {
		return apperrors.NotFound("enrollment not found", nil)
	}

	return nil
}

func (r *PostgresEnrollmentRepo) GetByStudentID(ctx context.Context, studentID int) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	query := `
		SELECT c.id, c.title, c.description, c.price, c.level, c.is_active, c.teacher_id, c.created_at, c.updated_at
		FROM courses c
		JOIN enrollments e ON e.course_id = c.id
		WHERE e.student_id = $1 AND e.deleted_at IS NULL AND c.deleted_at IS NULL
		ORDER BY e.created_at DESC
	`

	err := r.db.SelectContext(ctx, &courses, query, studentID)
	if err != nil {
		return nil, apperrors.Internal("failed to get enrolled courses", err)
	}

	return courses, nil
}

func (r *PostgresEnrollmentRepo) DeleteByCourseID(ctx context.Context, tx *sqlx.Tx, courseID int) error {
	query := `
		UPDATE enrollments
		SET deleted_at = NOW(),
		    updated_at = NOW()
		WHERE course_id = $1
			AND deleted_at IS NULL
	`

	_, err := tx.ExecContext(ctx, query, courseID)
	if err != nil {
		return apperrors.Internal("failed to delete enrollments for course", err)
	}

	return nil
}
