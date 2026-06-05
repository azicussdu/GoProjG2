package repository

import (
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
)

type PostgresCourseRepo struct {
	coursesMap map[int]model.Course
}

func NewPostgresCourseRepo() *PostgresCourseRepo {
	repo := &PostgresCourseRepo{
		coursesMap: make(map[int]model.Course),
	}

	repo.coursesMap[1] = model.Course{ID: 1, Title: "Golang 1", Price: 40000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.coursesMap[2] = model.Course{ID: 2, Title: "Python", Price: 55000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.coursesMap[3] = model.Course{ID: 3, Title: "Data Science", Price: 35000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	return repo
}

func (cr *PostgresCourseRepo) GetAll() ([]model.Course, error) {
	coursesSlice := make([]model.Course, 0)

	for _, crs := range cr.coursesMap {
		coursesSlice = append(coursesSlice, crs)
	}

	return coursesSlice, nil
}

func (cr *PostgresCourseRepo) GetByID(id int) (model.Course, error) {
	course, ok := cr.coursesMap[id]
	if !ok {
		return model.Course{}, apperrors.NotFound("course is not found", nil)
	}
	return course, nil
}

func (cr *PostgresCourseRepo) Delete(id int) error {
	if _, ok := cr.coursesMap[id]; !ok {
		return apperrors.NotFound("course is not found", nil)
	}

	delete(cr.coursesMap, id)
	return nil
}

func (cr *PostgresCourseRepo) Create(course model.Course) (int, error) {
	course.ID = len(cr.coursesMap) + 1
	cr.coursesMap[course.ID] = course

	return course.ID, nil
}

func (cr *PostgresCourseRepo) Update(id int, input model.UpdateCourse) (model.Course, error) {
	course, ok := cr.coursesMap[id]
	if !ok {
		return model.Course{}, apperrors.NotFound("course is not found", nil)
	}

	if input.Title != nil {
		course.Title = *input.Title
	}
	if input.Price != nil {
		course.Price = *input.Price
	}
	if input.IsActive != nil {
		course.IsActive = *input.IsActive
	}
	course.UpdatedAt = time.Now()

	cr.coursesMap[id] = course
	return course, nil
}
