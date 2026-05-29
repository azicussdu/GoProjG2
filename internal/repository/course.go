package repository

import (
	"errors"
	"time"

	"github.com/azicussdu/GoProjG2/internal/model"
)

type CourseRepo struct {
	coursesMap map[int]model.Course
}

func NewCourseRepo() *CourseRepo {
	repo := &CourseRepo{
		coursesMap: make(map[int]model.Course),
	}

	repo.coursesMap[1] = model.Course{ID: 1, Title: "Golang 1", Price: 40000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.coursesMap[2] = model.Course{ID: 2, Title: "Python", Price: 55000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.coursesMap[3] = model.Course{ID: 3, Title: "Data Science", Price: 35000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	return repo
}

func (cr *CourseRepo) GetAll() ([]model.Course, error) {
	coursesSlice := make([]model.Course, 0)

	for _, crs := range cr.coursesMap {
		coursesSlice = append(coursesSlice, crs)
	}

	return coursesSlice, nil
}

func (cr *CourseRepo) GetByID(id int) (model.Course, error) {
	course, ok := cr.coursesMap[id]
	if !ok {
		return model.Course{}, errors.New("course is not found")
	}
	return course, nil
}

func (cr *CourseRepo) Delete(id int) error {
	if _, ok := cr.coursesMap[id]; !ok {
		return errors.New("Course is not found")
	}

	delete(cr.coursesMap, id)
	return nil
}
