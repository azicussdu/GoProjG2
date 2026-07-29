package repository

import (
	"context"
	"errors"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"gorm.io/gorm"
)

type PostgresCourseRepo struct {
	db *gorm.DB
}

func NewPostgresCourseRepo(dbObj *gorm.DB) *PostgresCourseRepo {
	repo := &PostgresCourseRepo{
		db: dbObj,
	}

	return repo
}

func (pcr *PostgresCourseRepo) GetAll(ctx context.Context) ([]model.Course, error) {
	coursesSlice := make([]model.Course, 0) // Course - courses

	err := pcr.db.WithContext(ctx).Order("created_at").Find(&coursesSlice).Error
	if err != nil {
		return coursesSlice, apperrors.Internal("get all courses", err)
	}

	return coursesSlice, nil
}

func (pcr *PostgresCourseRepo) GetByID(ctx context.Context, id int) (model.Course, error) {
	var course model.Course

	err := pcr.db.WithContext(ctx).First(&course, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Course{}, apperrors.NotFound("course with ID not found", err)
		}
		return model.Course{}, apperrors.Internal("failed to get course", err)
	}

	return course, nil
}

func (pcr *PostgresCourseRepo) Delete(ctx context.Context, tx *gorm.DB, id int) error {
	result := tx.WithContext(ctx).Delete(&model.Course{}, id)
	if result.Error != nil {
		return apperrors.Internal("failed to delete course", result.Error)
	}

	if result.RowsAffected == 0 {
		return apperrors.NotFound("course with ID not found", nil)
	}

	return nil
}

func (pcr *PostgresCourseRepo) Create(ctx context.Context, course model.Course) (int, error) {
	err := pcr.db.WithContext(ctx).Omit("Teacher").Create(&course).Error
	if err != nil {
		return 0, apperrors.Internal("failed to created course", err)
	}

	return course.ID, nil
}

func (pcr *PostgresCourseRepo) Update(ctx context.Context, id int, input model.UpdateCourse) (int, error) {
	updates := map[string]any{}

	if input.Title != nil {
		updates["title"] = *input.Title
	}

	if input.Description != nil {
		updates["description"] = *input.Description
	}

	if input.Price != nil {
		updates["price"] = *input.Price
	}

	if input.Level != nil {
		updates["level"] = *input.Level
	}

	if input.IsActive != nil {
		updates["is_active"] = *input.IsActive
	}

	if input.TeacherID != nil {
		updates["teacher_id"] = *input.TeacherID
	}

	if len(updates) == 0 {
		return 0, apperrors.BadRequest("no fields to update", nil)
	}

	updates["updated_at"] = time.Now()

	result := pcr.db.WithContext(ctx).Model(&model.Course{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return 0, apperrors.Internal("failed to update course", result.Error)
	}

	if result.RowsAffected == 0 {
		return 0, apperrors.NotFound("course with ID not found", nil)
	}

	return id, nil
}
