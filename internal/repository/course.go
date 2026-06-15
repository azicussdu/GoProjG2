package repository

import (
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jmoiron/sqlx"
)

type PostgresCourseRepo struct {
	db *sqlx.DB
}

func NewPostgresCourseRepo(dbObj *sqlx.DB) *PostgresCourseRepo {
	repo := &PostgresCourseRepo{
		db: dbObj,
	}

	return repo
}

func (pcr *PostgresCourseRepo) GetAll() ([]model.Course, error) {
	coursesSlice := make([]model.Course, 0)

	query := `
		SELECT id, title, price, is_active
		FROM courses
		ORDER BY created_at
	`

	err := pcr.db.Select(&coursesSlice, query)
	if err != nil {
		return nil, err
	}

	return coursesSlice, nil
}

func (pcr *PostgresCourseRepo) GetByID(id int) (model.Course, error) {
	var course model.Course

	query := `
		SELECT id, title, price, is_active
		FROM courses
		WHERE id = $1
		LIMIT 1
	`

	err := pcr.db.Get(&course, query, id)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (pcr *PostgresCourseRepo) Delete(id int) error {
	query := `
		DELETE FROM courses
		WHERE id = $1
	`

	_, err := pcr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (pcr *PostgresCourseRepo) Create(course model.Course) (int, error) {
	return 0, nil
}

func (pcr *PostgresCourseRepo) Update(id int, input model.UpdateCourse) (model.Course, error) {
	//course, ok := cr.coursesMap[id]
	//if !ok {
	//	return model.Course{}, apperrors.NotFound("course is not found", nil)
	//}
	//
	//if input.Title != nil {
	//	course.Title = *input.Title
	//}
	//if input.Price != nil {
	//	course.Price = *input.Price
	//}
	//if input.IsActive != nil {
	//	course.IsActive = *input.IsActive
	//}
	//course.UpdatedAt = time.Now()
	//
	//cr.coursesMap[id] = course
	return model.Course{}, nil
}
