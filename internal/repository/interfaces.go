package repository

import (
	"context"

	"github.com/azicussdu/GoProjG2/internal/model"
)

type CourseRepoI interface {
	GetAll(ctx context.Context) ([]model.Course, error)
	GetByID(ctx context.Context, id int) (model.Course, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, course model.Course) (int, error)
	Update(ctx context.Context, id int, input model.UpdateCourse) (int, error)
}
