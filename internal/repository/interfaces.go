package repository

import "github.com/azicussdu/GoProjG2/internal/model"

type CourseRepoI interface {
	GetAll() ([]model.Course, error)
	GetByID(id int) (model.Course, error)
	Delete(id int) error
	Create(course model.Course) (int, error)
	Update(id int, input model.UpdateCourse) (int, error)
}
