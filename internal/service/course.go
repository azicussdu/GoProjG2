package service

import (
	"errors"

	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepo
}

func NewCourseService(courseRepo *repository.CourseRepo) *CourseService {
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
		return errors.New("Invalid ID")
	}
	return cs.repo.Delete(id)
}
