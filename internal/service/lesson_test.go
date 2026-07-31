package service

import (
	"context"
	"net/http"
	"testing"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLessonService_GetByCourseID(t *testing.T) {
	t.Run("invalid course ID", func(t *testing.T) {
		svc := NewLessonService(new(MockLessonRepo))

		_, err := svc.GetByCourseID(context.Background(), 0)

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("delegates to repo", func(t *testing.T) {
		want := []model.Lesson{{ID: 1, Title: "Intro"}}
		repo := new(MockLessonRepo)
		repo.On("GetByCourseID", mock.Anything, 5).Return(want, nil)
		svc := NewLessonService(repo)

		got, err := svc.GetByCourseID(context.Background(), 5)

		require.NoError(t, err)
		require.Len(t, got, 1)
		require.Equal(t, want[0].ID, got[0].ID)
		repo.AssertExpectations(t)
	})
}

func TestLessonService_Create(t *testing.T) {
	t.Run("invalid input is rejected before hitting the repo", func(t *testing.T) {
		svc := NewLessonService(new(MockLessonRepo))

		_, err := svc.Create(context.Background(), model.CreateLesson{CourseID: 1, Title: "Hi"})

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("defaults position to 1 when not set", func(t *testing.T) {
		repo := new(MockLessonRepo)
		repo.On("Create", mock.Anything, mock.MatchedBy(func(l model.Lesson) bool {
			return l.Position == 1
		})).Return(1, nil)
		svc := NewLessonService(repo)

		_, err := svc.Create(context.Background(), model.CreateLesson{CourseID: 1, Title: "Intro", Position: 0})

		require.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("keeps an explicit position", func(t *testing.T) {
		repo := new(MockLessonRepo)
		repo.On("Create", mock.Anything, mock.MatchedBy(func(l model.Lesson) bool {
			return l.Position == 3
		})).Return(1, nil)
		svc := NewLessonService(repo)

		_, err := svc.Create(context.Background(), model.CreateLesson{CourseID: 1, Title: "Intro", Position: 3})

		require.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLessonService_Update(t *testing.T) {
	t.Run("invalid ID", func(t *testing.T) {
		svc := NewLessonService(new(MockLessonRepo))

		_, err := svc.Update(context.Background(), 0, model.UpdateLesson{})

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("propagates not-found from the existence check", func(t *testing.T) {
		repo := new(MockLessonRepo)
		repo.On("GetByID", mock.Anything, 99).Return(model.Lesson{}, apperrors.NotFound("lesson not found", nil))
		svc := NewLessonService(repo)

		_, err := svc.Update(context.Background(), 99, model.UpdateLesson{})

		require.Equal(t, http.StatusNotFound, appErrorCode(err))
		repo.AssertExpectations(t)
	})

	t.Run("success calls repo.Update", func(t *testing.T) {
		repo := new(MockLessonRepo)
		repo.On("GetByID", mock.Anything, 5).Return(model.Lesson{ID: 5}, nil)
		repo.On("Update", mock.Anything, 5, mock.Anything).Return(5, nil)
		svc := NewLessonService(repo)

		newTitle := "Updated Title"
		id, err := svc.Update(context.Background(), 5, model.UpdateLesson{Title: &newTitle})

		require.NoError(t, err)
		require.Equal(t, 5, id)
		repo.AssertExpectations(t)
	})
}

func TestLessonService_Delete(t *testing.T) {
	t.Run("invalid ID", func(t *testing.T) {
		svc := NewLessonService(new(MockLessonRepo))

		err := svc.Delete(context.Background(), 0)

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("delegates to repo", func(t *testing.T) {
		repo := new(MockLessonRepo)
		repo.On("Delete", mock.Anything, 5).Return(nil)
		svc := NewLessonService(repo)

		require.NoError(t, svc.Delete(context.Background(), 5))
		repo.AssertExpectations(t)
	})
}
