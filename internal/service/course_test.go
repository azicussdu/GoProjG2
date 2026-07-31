package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Note: CourseService.Delete opens a real *sqlx.Tx via cs.db.BeginTxx,
// so it can't be exercised with these mocks. That method is a better
// fit for an integration test against a real (or sqlmock) database -
// not every method has to be a unit test.

func TestCourseService_GetAll(t *testing.T) {
	t.Run("propagates repo error", func(t *testing.T) {
		repo := new(MockCourseRepo)
		repo.On("GetAll", mock.Anything).Return(nil, errors.New("db is down"))
		svc := NewCourseService(repo, new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		_, err := svc.GetAll(context.Background())

		require.Error(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("returns courses from repo", func(t *testing.T) {
		want := []model.Course{{ID: 1, Title: "Go Basics"}}
		repo := new(MockCourseRepo)
		repo.On("GetAll", mock.Anything).Return(want, nil)
		svc := NewCourseService(repo, new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		got, err := svc.GetAll(context.Background())

		require.NoError(t, err)
		require.Len(t, got, 1)
		assert.Equal(t, want[0].ID, got[0].ID)
		repo.AssertExpectations(t)
	})
}

func TestCourseService_GetByID(t *testing.T) {
	t.Run("invalid ID", func(t *testing.T) {
		svc := NewCourseService(new(MockCourseRepo), new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		_, err := svc.GetByID(context.Background(), 0)

		var appErr *apperrors.AppError
		require.True(t, errors.As(err, &appErr))
		assert.Equal(t, http.StatusBadRequest, appErr.Code)
	})

	t.Run("delegates to repo", func(t *testing.T) {
		repo := new(MockCourseRepo)
		repo.On("GetByID", mock.Anything, 5).Return(model.Course{ID: 5, Title: "Go Basics"}, nil)
		svc := NewCourseService(repo, new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		got, err := svc.GetByID(context.Background(), 5)

		require.NoError(t, err)
		assert.Equal(t, 5, got.ID)
		repo.AssertExpectations(t)
	})
}

func TestCourseService_Create(t *testing.T) {
	t.Run("invalid input is rejected before hitting the repo", func(t *testing.T) {
		repo := new(MockCourseRepo)
		svc := NewCourseService(repo, new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		_, err := svc.Create(context.Background(), model.CreateCourse{Title: "Hi", TeacherID: 1})

		var appErr *apperrors.AppError
		require.True(t, errors.As(err, &appErr))
		assert.Equal(t, http.StatusBadRequest, appErr.Code)
		repo.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run("success calls repo.Create", func(t *testing.T) {
		repo := new(MockCourseRepo)
		repo.On("Create", mock.Anything, mock.MatchedBy(func(c model.Course) bool {
			return c.Title == "Go Basics"
		})).Return(10, nil)
		svc := NewCourseService(repo, new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		id, err := svc.Create(context.Background(), model.CreateCourse{Title: "Go Basics", TeacherID: 1})

		require.NoError(t, err)
		assert.Equal(t, 10, id)
		repo.AssertExpectations(t)
	})
}

func TestCourseService_Update(t *testing.T) {
	t.Run("invalid ID", func(t *testing.T) {
		svc := NewCourseService(new(MockCourseRepo), new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		_, err := svc.Update(context.Background(), 0, model.UpdateCourse{})

		var appErr *apperrors.AppError
		require.True(t, errors.As(err, &appErr))
		assert.Equal(t, http.StatusBadRequest, appErr.Code)
	})

	t.Run("propagates not-found from the existence check", func(t *testing.T) {
		repo := new(MockCourseRepo)
		repo.On("GetByID", mock.Anything, 99).Return(model.Course{}, apperrors.NotFound("course with ID not found", nil))
		svc := NewCourseService(repo, new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		_, err := svc.Update(context.Background(), 99, model.UpdateCourse{})

		var appErr *apperrors.AppError
		require.True(t, errors.As(err, &appErr))
		assert.Equal(t, http.StatusNotFound, appErr.Code)
		repo.AssertExpectations(t)
	})

	t.Run("success calls repo.Update", func(t *testing.T) {
		repo := new(MockCourseRepo)
		repo.On("GetByID", mock.Anything, 5).Return(model.Course{ID: 5}, nil)
		repo.On("Update", mock.Anything, 5, mock.Anything).Return(5, nil)
		svc := NewCourseService(repo, new(MockLessonRepo), new(MockEnrollmentRepo), nil)

		newTitle := "Updated Title"
		id, err := svc.Update(context.Background(), 5, model.UpdateCourse{Title: &newTitle})

		require.NoError(t, err)
		assert.Equal(t, 5, id)
		repo.AssertExpectations(t)
	})
}
