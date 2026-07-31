package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// appErrorCode is a small test helper: it unwraps an error into an
// *apperrors.AppError (the same way internal/handler/error_response.go
// does) and returns its HTTP status code, or 0 if err isn't an AppError.
func appErrorCode(err error) int {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return appErr.Code
	}
	return 0
}

func TestEnrollmentService_Enroll_InvalidStudentID(t *testing.T) {
	svc := NewEnrollmentService(new(MockEnrollmentRepo), new(MockCourseRepo))

	_, err := svc.Enroll(context.Background(), 0, 5)

	require.Equal(t, http.StatusBadRequest, appErrorCode(err))
}

func TestEnrollmentService_Enroll_CourseNotFound(t *testing.T) {
	courseRepo := new(MockCourseRepo)
	courseRepo.On("GetByID", mock.Anything, 999).Return(model.Course{}, apperrors.NotFound("course with ID not found", nil))
	svc := NewEnrollmentService(new(MockEnrollmentRepo), courseRepo)

	_, err := svc.Enroll(context.Background(), 1, 999)

	require.Equal(t, http.StatusNotFound, appErrorCode(err))
	courseRepo.AssertExpectations(t)
}

func TestEnrollmentService_Enroll_InactiveCourse(t *testing.T) {
	courseRepo := new(MockCourseRepo)
	courseRepo.On("GetByID", mock.Anything, 5).Return(model.Course{ID: 5, IsActive: false}, nil)
	svc := NewEnrollmentService(new(MockEnrollmentRepo), courseRepo)

	_, err := svc.Enroll(context.Background(), 1, 5)

	require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	courseRepo.AssertExpectations(t)
}

func TestEnrollmentService_Enroll_Success(t *testing.T) {
	courseRepo := new(MockCourseRepo)
	courseRepo.On("GetByID", mock.Anything, 5).Return(model.Course{ID: 5, IsActive: true}, nil)

	enrollmentRepo := new(MockEnrollmentRepo)
	enrollmentRepo.On("Create", mock.Anything, mock.MatchedBy(func(e model.Enrollment) bool {
		return e.StudentID == 7 && e.CourseID == 5
	})).Return(42, nil)

	svc := NewEnrollmentService(enrollmentRepo, courseRepo)

	id, err := svc.Enroll(context.Background(), 7, 5)

	require.NoError(t, err)
	require.Equal(t, 42, id)
	courseRepo.AssertExpectations(t)
	enrollmentRepo.AssertExpectations(t)
}

func TestEnrollmentService_Leave(t *testing.T) {
	t.Run("invalid course ID", func(t *testing.T) {
		svc := NewEnrollmentService(new(MockEnrollmentRepo), new(MockCourseRepo))

		err := svc.Leave(context.Background(), 1, 0)

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("delegates to repo on valid IDs", func(t *testing.T) {
		enrollmentRepo := new(MockEnrollmentRepo)
		enrollmentRepo.On("Delete", mock.Anything, 7, 5).Return(nil)
		svc := NewEnrollmentService(enrollmentRepo, new(MockCourseRepo))

		require.NoError(t, svc.Leave(context.Background(), 7, 5))
		enrollmentRepo.AssertExpectations(t)
	})
}

func TestEnrollmentService_GetMyCourses(t *testing.T) {
	t.Run("invalid student ID", func(t *testing.T) {
		svc := NewEnrollmentService(new(MockEnrollmentRepo), new(MockCourseRepo))

		_, err := svc.GetMyCourses(context.Background(), -1)

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("returns courses from repo", func(t *testing.T) {
		want := []model.Course{{ID: 1, Title: "Go Basics"}}
		enrollmentRepo := new(MockEnrollmentRepo)
		enrollmentRepo.On("GetByStudentID", mock.Anything, 7).Return(want, nil)
		svc := NewEnrollmentService(enrollmentRepo, new(MockCourseRepo))

		got, err := svc.GetMyCourses(context.Background(), 7)

		require.NoError(t, err)
		require.Len(t, got, 1)
		require.Equal(t, want[0].ID, got[0].ID)
		enrollmentRepo.AssertExpectations(t)
	})
}
