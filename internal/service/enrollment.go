package service

import (
	"context"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/repository"
)

type EnrollmentService struct {
	repo       repository.EnrollmentRepoI
	courseRepo repository.CourseRepoI
}

func NewEnrollmentService(enrollmentRepo repository.EnrollmentRepoI, courseRepo repository.CourseRepoI) *EnrollmentService {
	return &EnrollmentService{repo: enrollmentRepo, courseRepo: courseRepo}
}

func (s *EnrollmentService) Enroll(ctx context.Context, studentID, courseID int) (int, error) {
	if studentID <= 0 {
		return 0, apperrors.BadRequest("invalid student ID", nil)
	}
	if courseID <= 0 {
		return 0, apperrors.BadRequest("invalid course ID parameter", nil)
	}

	course, err := s.courseRepo.GetByID(ctx, courseID)
	if err != nil {
		return 0, apperrors.NotFound("course not found", nil)
	}

	if !course.IsActive {
		return 0, apperrors.BadRequest("course is not active", nil)
	}

	enrollment := model.Enrollment{
		StudentID: studentID,
		CourseID:  courseID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.Create(ctx, enrollment)
}

func (s *EnrollmentService) Leave(ctx context.Context, studentID, courseID int) error {
	if studentID <= 0 {
		return apperrors.BadRequest("invalid student ID", nil)
	}
	if courseID <= 0 {
		return apperrors.BadRequest("invalid course ID parameter", nil)
	}

	return s.repo.Delete(ctx, studentID, courseID)
}

func (s *EnrollmentService) GetMyCourses(ctx context.Context, studentID int) ([]model.Course, error) {
	if studentID <= 0 {
		return nil, apperrors.BadRequest("invalid student ID", nil)
	}

	return s.repo.GetByStudentID(ctx, studentID)
}
