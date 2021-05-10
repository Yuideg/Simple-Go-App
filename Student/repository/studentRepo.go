package repository
import (
	"fmt"
	"github.com/Yuideg/firstApp/entity"
	"github.com/jinzhu/gorm"
)
// RoomRepositoryImpl implements the rooms.RoomRepository interface
type GormRepositoryImpl struct {
	Con *gorm.DB
}

// NewRoomRepositoryImpl will create an object of PsqlRoomRepository
func NewGormRepositoryImpl(Conn *gorm.DB) *GormRepositoryImpl {
	return &GormRepositoryImpl{Con: Conn}
}

// News returns all rooms from the database
func (rri *GormRepositoryImpl) Students() ([]entity.StudentInfo, []error) {
	cmnts := []entity.StudentInfo{}
	fmt.Println("ents gormrepo")
	errs := rri.Con.Find(&cmnts).GetErrors()
	fmt.Println(cmnts)
	if len(errs) > 0 {
		return nil, nil
	}
	return cmnts, nil
}

// NewsById returns a News with a given id
func (rri *GormRepositoryImpl) Student(id int) (*entity.StudentInfo, []error) {
	cmnt := &entity.StudentInfo{}
	errs := rri.Con.First(cmnt, id).GetErrors()
	if len(errs) > 0 {
		fmt.Println("inside evid errr")
		return nil, errs
	}
	return cmnt, nil
}

// UpdateNews updates a given object with a new data
func (rri *GormRepositoryImpl) UpdateStudentInfor(r *entity.StudentInfo) (*entity.StudentInfo, []error) {
	cmnt := r
	errs := rri.Con.Save(&cmnt).GetErrors()
	if errs != nil {
		return nil, errs
	}
	return nil, nil
}

// DeleteNews removes a News from a database by its id
func (rri *GormRepositoryImpl) DeleteStudent(id int) (*entity.StudentInfo, []error) {
	cmnt, errs := rri.Student(id)

	if errs != nil {
		return nil, errs
	}
	errs = rri.Con.Delete(&cmnt, id).GetErrors()
	if errs != nil {
		return nil, errs
	}
	return cmnt, errs
}

// StoreNews stores new News information to database
func (rri *GormRepositoryImpl) RegisterStudent(r entity.StudentInfo) (*entity.StudentInfo, []error) {
	cmnt := r
	errs := rri.Con.Create(&cmnt).GetErrors()
	if errs != nil {
		return nil, errs
	}
	return nil, errs
}
