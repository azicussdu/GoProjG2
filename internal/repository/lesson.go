package repository

import (
	"context"
	"errors"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"gorm.io/gorm"
)

type PostgresLessonRepo struct {
	db *gorm.DB
}

func NewPostgresLessonRepo(db *gorm.DB) *PostgresLessonRepo {
	return &PostgresLessonRepo{db: db}
}

func (r *PostgresLessonRepo) GetByCourseID(ctx context.Context, courseID int) ([]model.Lesson, error) {
	lessons := make([]model.Lesson, 0)

	err := r.db.WithContext(ctx).Where("course_id = ?", courseID).Order("position").Find(&lessons).Error
	if err != nil {
		return nil, apperrors.Internal("failed to get lessons", err)
	}

	return lessons, nil
}

func (r *PostgresLessonRepo) GetByID(ctx context.Context, id int) (model.Lesson, error) {
	var lesson model.Lesson

	err := r.db.WithContext(ctx).First(&lesson, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Lesson{}, apperrors.NotFound("lesson with ID not found", err)
		}
		return model.Lesson{}, apperrors.Internal("failed to get lesson", err)
	}

	return lesson, nil
}

func (r *PostgresLessonRepo) Create(ctx context.Context, lesson model.Lesson) (int, error) {
	err := r.db.WithContext(ctx).Create(&lesson).Error
	if err != nil {
		return 0, apperrors.Internal("failed to create lesson", err)
	}

	return lesson.ID, nil
}

func (r *PostgresLessonRepo) Update(ctx context.Context, id int, input model.UpdateLesson) (int, error) {
	updates := map[string]any{}

	if input.Title != nil {
		updates["title"] = *input.Title
	}

	if input.Content != nil {
		updates["content"] = *input.Content
	}

	if input.Position != nil {
		updates["position"] = *input.Position
	}

	if len(updates) == 0 {
		return 0, apperrors.BadRequest("no fields to update", nil)
	}

	updates["updated_at"] = time.Now()

	result := r.db.WithContext(ctx).Model(&model.Lesson{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return 0, apperrors.Internal("failed to update lesson", result.Error)
	}

	if result.RowsAffected == 0 {
		return 0, apperrors.NotFound("lesson with ID not found", nil)
	}

	return id, nil
}

func (r *PostgresLessonRepo) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&model.Lesson{}, id)
	if result.Error != nil {
		return apperrors.Internal("failed to delete lesson", result.Error)
	}

	if result.RowsAffected == 0 {
		return apperrors.NotFound("lesson with ID not found", nil)
	}

	return nil
}

func (r *PostgresLessonRepo) DeleteByCourseID(ctx context.Context, tx *gorm.DB, courseID int) error {
	err := tx.WithContext(ctx).Where("course_id = ?", courseID).Delete(&model.Lesson{}).Error
	if err != nil {
		return apperrors.Internal("failed to delete lessons", err)
	}

	return nil
}
