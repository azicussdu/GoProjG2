package service

import (
	"context"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/repository"
	"github.com/jmoiron/sqlx"
)

type CourseService struct {
	repo           repository.CourseRepoI
	lessonRepo     repository.LessonRepoI
	enrollmentRepo repository.EnrollmentRepoI
	db             *sqlx.DB
}

func NewCourseService(
	courseRepo repository.CourseRepoI,
	lessonRepo repository.LessonRepoI,
	enrollmentRepo repository.EnrollmentRepoI,
	db *sqlx.DB) *CourseService {

	service := &CourseService{
		repo:           courseRepo,
		lessonRepo:     lessonRepo,
		enrollmentRepo: enrollmentRepo,
		db:             db,
	}

	return service
}

func (cs *CourseService) GetAll(ctx context.Context) ([]model.Course, error) {
	courses, err := cs.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (cs *CourseService) GetByID(ctx context.Context, id int) (model.Course, error) {
	if id <= 0 {
		return model.Course{}, apperrors.BadRequest("invalid ID parameter", nil)
	}
	return cs.repo.GetByID(ctx, id)
}

func (cs *CourseService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return apperrors.BadRequest("invalid ID parameter", nil)
	}

	tx, err := cs.db.BeginTxx(ctx, nil)
	if err != nil {
		return apperrors.Internal("failed to begin transaction", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	if err = cs.lessonRepo.DeleteByCourseID(ctx, tx, id); err != nil {
		return err
	}
	if err = cs.enrollmentRepo.DeleteByCourseID(ctx, tx, id); err != nil {
		return err
	}
	if err = cs.repo.Delete(ctx, tx, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return apperrors.Internal("failed to commit course delete", err)
	}
	return nil
}

func (cs *CourseService) Create(ctx context.Context, input model.CreateCourse) (int, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	course := model.Course{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Level:       input.Level,
		IsActive:    input.IsActive,
		TeacherID:   input.TeacherID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return cs.repo.Create(ctx, course)
}

func (cs *CourseService) Update(ctx context.Context, id int, input model.UpdateCourse) (int, error) {
	if id <= 0 {
		return 0, apperrors.BadRequest("invalid ID parameter", nil)
	}

	_, err := cs.repo.GetByID(ctx, id)
	if err != nil {
		return 0, err
	}

	if err = input.Validate(); err != nil {
		return 0, err
	}

	return cs.repo.Update(ctx, id, input)
}
