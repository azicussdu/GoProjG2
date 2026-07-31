package service

import (
	"context"

	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

// This file holds shared testify/mock mocks used by the service tests
// below. Each mock implements a repository interface (see
// internal/repository/interfaces.go) via mock.Mock, so tests set up
// expectations with mock.On(...).Return(...) instead of a real
// database - no Postgres, no network.
//
// A call that wasn't stubbed with mock.On panics, which makes it
// obvious in a test failure if a service calls a repo method the
// test didn't expect.

type MockCourseRepo struct {
	mock.Mock
}

func (m *MockCourseRepo) GetAll(ctx context.Context) ([]model.Course, error) {
	args := m.Called(ctx)
	var courses []model.Course
	if v := args.Get(0); v != nil {
		courses = v.([]model.Course)
	}
	return courses, args.Error(1)
}

func (m *MockCourseRepo) GetByID(ctx context.Context, id int) (model.Course, error) {
	args := m.Called(ctx, id)
	var course model.Course
	if v := args.Get(0); v != nil {
		course = v.(model.Course)
	}
	return course, args.Error(1)
}

func (m *MockCourseRepo) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	args := m.Called(ctx, tx, id)
	return args.Error(0)
}

func (m *MockCourseRepo) Create(ctx context.Context, course model.Course) (int, error) {
	args := m.Called(ctx, course)
	return args.Int(0), args.Error(1)
}

func (m *MockCourseRepo) Update(ctx context.Context, id int, input model.UpdateCourse) (int, error) {
	args := m.Called(ctx, id, input)
	return args.Int(0), args.Error(1)
}

type MockLessonRepo struct {
	mock.Mock
}

func (m *MockLessonRepo) GetByCourseID(ctx context.Context, courseID int) ([]model.Lesson, error) {
	args := m.Called(ctx, courseID)
	var lessons []model.Lesson
	if v := args.Get(0); v != nil {
		lessons = v.([]model.Lesson)
	}
	return lessons, args.Error(1)
}

func (m *MockLessonRepo) GetByID(ctx context.Context, id int) (model.Lesson, error) {
	args := m.Called(ctx, id)
	var lesson model.Lesson
	if v := args.Get(0); v != nil {
		lesson = v.(model.Lesson)
	}
	return lesson, args.Error(1)
}

func (m *MockLessonRepo) Create(ctx context.Context, lesson model.Lesson) (int, error) {
	args := m.Called(ctx, lesson)
	return args.Int(0), args.Error(1)
}

func (m *MockLessonRepo) Update(ctx context.Context, id int, input model.UpdateLesson) (int, error) {
	args := m.Called(ctx, id, input)
	return args.Int(0), args.Error(1)
}

func (m *MockLessonRepo) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockLessonRepo) DeleteByCourseID(ctx context.Context, tx *sqlx.Tx, courseID int) error {
	args := m.Called(ctx, tx, courseID)
	return args.Error(0)
}

type MockEnrollmentRepo struct {
	mock.Mock
}

func (m *MockEnrollmentRepo) Create(ctx context.Context, enrollment model.Enrollment) (int, error) {
	args := m.Called(ctx, enrollment)
	return args.Int(0), args.Error(1)
}

func (m *MockEnrollmentRepo) Delete(ctx context.Context, studentID, courseID int) error {
	args := m.Called(ctx, studentID, courseID)
	return args.Error(0)
}

func (m *MockEnrollmentRepo) GetByStudentID(ctx context.Context, studentID int) ([]model.Course, error) {
	args := m.Called(ctx, studentID)
	var courses []model.Course
	if v := args.Get(0); v != nil {
		courses = v.([]model.Course)
	}
	return courses, args.Error(1)
}

func (m *MockEnrollmentRepo) DeleteByCourseID(ctx context.Context, tx *sqlx.Tx, courseID int) error {
	args := m.Called(ctx, tx, courseID)
	return args.Error(0)
}

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(ctx context.Context, user model.User) (int, error) {
	args := m.Called(ctx, user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	args := m.Called(ctx, email)
	var user model.User
	if v := args.Get(0); v != nil {
		user = v.(model.User)
	}
	return user, args.Error(1)
}

func (m *MockUserRepo) GetByID(ctx context.Context, id int) (model.User, error) {
	args := m.Called(ctx, id)
	var user model.User
	if v := args.Get(0); v != nil {
		user = v.(model.User)
	}
	return user, args.Error(1)
}
