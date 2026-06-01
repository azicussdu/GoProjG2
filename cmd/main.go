package main

//import (
//	"fmt"
//
//	"github.com/azicussdu/GoProjG2/internal/model"
//)
//
//type CourseService struct {
//	repo CourseRepoI // (CourseRepoI = &CourseRepo{})
//}
//
//// --------------------------------------
//
//type CourseRepoI interface {
//	GetAll() ([]model.Course, error)
//	// GetByID, Create, Update, Delete
//}
//
//// --------------------------------------
//
//type CourseRepo struct {
//}
//
//func (cr *CourseRepo) GetAll() ([]model.Course, error) {
//	// SELECT * from courses
//	return nil, nil
//}
//
//// --------------------------------------
//
//type CourseMongoRepo struct {
//}
//
//func (cr *CourseMongoRepo) GetAll() ([]model.Course, error) {
//	// document.getCourses()
//	return nil, nil
//}
//
//// --------------------------------------
//
//func main() {
//	courseRepo := &CourseMongoRepo{}
//	service := CourseService{repo: courseRepo} // CourseRepoI = &CourseRepo{}
//
//	fmt.Println(service)
//}
