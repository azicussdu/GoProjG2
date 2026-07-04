package service

import (
	"context"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/repository"
)

type LessonService struct {
	repo repository.LessonRepoI
}

func NewLessonService(lessonRepo repository.LessonRepoI) *LessonService {
	return &LessonService{repo: lessonRepo}
}

func (s *LessonService) GetByCourseID(ctx context.Context, courseID int) ([]model.Lesson, error) {
	if courseID <= 0 {
		return nil, apperrors.BadRequest("invalid course ID parameter", nil)
	}
	return s.repo.GetByCourseID(ctx, courseID)
}

func (s *LessonService) GetByID(ctx context.Context, id int) (model.Lesson, error) {
	if id <= 0 {
		return model.Lesson{}, apperrors.BadRequest("invalid ID parameter", nil)
	}
	return s.repo.GetByID(ctx, id)
}

func (s *LessonService) Create(ctx context.Context, input model.CreateLesson) (int, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	position := input.Position
	if position == 0 {
		position = 1
	}

	lesson := model.Lesson{
		CourseID:  input.CourseID,
		Title:     input.Title,
		Content:   input.Content,
		Position:  position,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.Create(ctx, lesson)
}

func (s *LessonService) Update(ctx context.Context, id int, input model.UpdateLesson) (int, error) {
	if id <= 0 {
		return 0, apperrors.BadRequest("invalid ID parameter", nil)
	}

	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return 0, err
	}

	if err = input.Validate(); err != nil {
		return 0, err
	}

	return s.repo.Update(ctx, id, input)
}

func (s *LessonService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return apperrors.BadRequest("invalid ID parameter", nil)
	}
	return s.repo.Delete(ctx, id)
}
