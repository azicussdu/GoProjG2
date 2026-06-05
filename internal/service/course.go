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
	return cs.repo.GetByID(id)
}

func (cs *CourseService) Delete(id int) error {
	if id <= 0 {
		return apperrors.BadRequest("invalid ID parameter", nil)
	}
	return cs.repo.Delete(id)
}

func (cs *CourseService) Create(input model.CreateCourse) (int, error) {
	if len(input.Title) < 3 {
		return 0, apperrors.BadRequest("course title is too short", nil)
	}
	if input.Price < 0 {
		return 0, apperrors.BadRequest("price can not be negative", nil)
	}

	course := model.Course{
		Title:     input.Title,
		Price:     input.Price,
		IsActive:  input.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return cs.repo.Create(course)
}

func (cs *CourseService) Update(id int, input model.UpdateCourse) (model.Course, error) {
	// kurs bar ma jok pa?
	_, err := cs.repo.GetByID(id)
	if err != nil {
		return model.Course{}, err
	}

	if input.Title != nil && len(*input.Title) < 3 {
		return model.Course{}, apperrors.BadRequest("course title is too short", nil)
	}

	if input.Price != nil && *input.Price < 0 {
		return model.Course{}, apperrors.BadRequest("price can not be negative", nil)
	}

	return cs.repo.Update(id, input)
}
