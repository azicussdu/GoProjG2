package repository

import (
	"context"
	"errors"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type PostgresEnrollmentRepo struct {
	db *gorm.DB
}

func NewPostgresEnrollmentRepo(db *gorm.DB) *PostgresEnrollmentRepo {
	return &PostgresEnrollmentRepo{db: db}
}

func (r *PostgresEnrollmentRepo) Create(ctx context.Context, enrollment model.Enrollment) (int, error) {
	err := r.db.WithContext(ctx).Create(&enrollment).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
			return 0, apperrors.Conflict("already enrolled in this course", err)
		}
		return 0, apperrors.Internal("failed to create enrollment", err)
	}

	return enrollment.ID, nil
}

func (r *PostgresEnrollmentRepo) Delete(ctx context.Context, studentID, courseID int) error {
	result := r.db.WithContext(ctx).
		Where("student_id = ? AND course_id = ?", studentID, courseID).
		Delete(&model.Enrollment{})
	if result.Error != nil {
		return apperrors.Internal("failed to leave course", result.Error)
	}

	if result.RowsAffected == 0 {
		return apperrors.NotFound("enrollment not found", nil)
	}

	return nil
}

func (r *PostgresEnrollmentRepo) GetByStudentID(ctx context.Context, studentID int) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	err := r.db.WithContext(ctx).
		Joins("JOIN enrollments ON enrollments.course_id = courses.id").
		Where("enrollments.student_id = ? AND enrollments.deleted_at IS NULL", studentID).
		Order("enrollments.created_at DESC").
		Find(&courses).Error
	if err != nil {
		return nil, apperrors.Internal("failed to get enrolled courses", err)
	}

	return courses, nil
}

func (r *PostgresEnrollmentRepo) DeleteByCourseID(ctx context.Context, tx *gorm.DB, courseID int) error {
	err := tx.WithContext(ctx).Where("course_id = ?", courseID).Delete(&model.Enrollment{}).Error
	if err != nil {
		return apperrors.Internal("failed to delete enrollments for course", err)
	}

	return nil
}
