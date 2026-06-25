package service

import (
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/repository"
)

type CourseService struct {
	repo repository.CourseRepoI
}

func NewCourseService(courseRepo repository.CourseRepoI) *CourseService {
	service := &CourseService{
		repo: courseRepo,
	}

	return service
}

func (cs *CourseService) GetAll() ([]model.Course, error) {
	courses, err := cs.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (cs *CourseService) GetByID(id int) (model.Course, error) {
	if id <= 0 {
		return model.Course{}, apperrors.BadRequest("invalid ID parameter", nil)
	}
	return cs.repo.GetByID(id)
}

func (cs *CourseService) Delete(id int) error {
	if id <= 0 {
		return apperrors.BadRequest("invalid ID parameter", nil)
	}
	return cs.repo.Delete(id)
}

func (cs *CourseService) Create(input model.CreateCourse) (int, error) {
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

	return cs.repo.Create(course)
}

func (cs *CourseService) Update(id int, input model.UpdateCourse) (int, error) {
	if id <= 0 {
		return 0, apperrors.BadRequest("invalid ID parameter", nil)
	}

	_, err := cs.repo.GetByID(id)
	if err != nil {
		return 0, err
	}

	if err = input.Validate(); err != nil {
		return 0, err
	}

	return cs.repo.Update(id, input)
}
