package repository

import (
	"context"

	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jmoiron/sqlx"
)

type CourseRepoI interface {
	GetAll(ctx context.Context) ([]model.Course, error)
	GetByID(ctx context.Context, id int) (model.Course, error)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
	Create(ctx context.Context, course model.Course) (int, error)
	Update(ctx context.Context, id int, input model.UpdateCourse) (int, error)
}

type LessonRepoI interface {
	GetByCourseID(ctx context.Context, courseID int) ([]model.Lesson, error)
	GetByID(ctx context.Context, id int) (model.Lesson, error)
	Create(ctx context.Context, lesson model.Lesson) (int, error)
	Update(ctx context.Context, id int, input model.UpdateLesson) (int, error)
	Delete(ctx context.Context, id int) error
	DeleteByCourseID(ctx context.Context, tx *sqlx.Tx, courseID int) error
}

type UserRepoI interface {
	Create(ctx context.Context, user model.User) (int, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetByID(ctx context.Context, id int) (model.User, error)
}

type EnrollmentRepoI interface {
	Create(ctx context.Context, enrollment model.Enrollment) (int, error)
	Delete(ctx context.Context, studentID, courseID int) error
	GetByStudentID(ctx context.Context, studentID int) ([]model.Course, error)
	DeleteByCourseID(ctx context.Context, tx *sqlx.Tx, courseID int) error
}
