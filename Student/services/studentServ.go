package services

import (
	"github.com/Yuideg/firstApp/Student"
	"github.com/Yuideg/firstApp/entity"
)

// RoomServiceImpl implements rooms.RoomService interface
type GormServiceImpl struct {
	sturepo Student.StudentRepository
}

// NewNewsServiceImpl will create new RoomService object
func NewGorServiceImpl(st Student.StudentRepository) *GormServiceImpl {
	return &GormServiceImpl{sturepo: st}
}

// News returns list of all rooms
func (rs *GormServiceImpl) Students() ([]entity.StudentInfo, []error) {

	news, err := rs.sturepo.Students()

	if err != nil {
		return nil, err
	}

	return news, nil
}

// StoreNews persists new room information
func (rs *GormServiceImpl) RegisterStudent(st entity.StudentInfo) (*entity.StudentInfo, []error) {

	r,err := rs.sturepo.RegisterStudent(st)

	if err != nil {
		return nil,err
	}

	return r,nil
}

// single student
func (rs *GormServiceImpl)Student(id int) (*entity.StudentInfo, []error) {

	r, err := rs.sturepo.Student(int(id))

	if err != nil {
		return r, err
	}

	return r, nil
}

// UpdateNews updates a cateogory with new data
func (rs *GormServiceImpl) UpdateStudentInfor(st *entity.StudentInfo) (*entity.StudentInfo, []error) {
	r, err := rs.sturepo.UpdateStudentInfor(st)
	if err != nil {
		return r, err
	}
	return r, nil
}
// DeleteNews delete a room by its id
func (rs *GormServiceImpl) DeleteStudent(id int) (*entity.StudentInfo, []error) {
	r,err := rs.sturepo.DeleteStudent(int(id))
	if err != nil {
		return r,err
	}
	return r,nil
}
