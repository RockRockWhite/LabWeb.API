package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type TeachersRepository struct {
	db *gorm.DB
}

var _teachersRepository *TeachersRepository

func init() {
	db := getDB()
	if err := db.AutoMigrate(&entities.Teacher{}); err != nil {
		utils.GetLogger().Fatal("Fatal migrate database %s : %s \n", "Teacher", err)
	}

	_teachersRepository = &TeachersRepository{db}
}

func GetTeachersRepository() *TeachersRepository {
	return _teachersRepository
}

// GetTeacher 从id获得教师
func (repository *TeachersRepository) GetTeacher(id uint) (*entities.Teacher, error) {
	var err error
	var teacher entities.Teacher
	if result := repository.db.First(&teacher, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get teacher id %v : %s", id, result.Error))
	}

	return &teacher, err
}

// GetTeachers 获得教师列表
func (repository *TeachersRepository) GetTeachers(limit int, offset int) ([]entities.Teacher, error) {
	var err error
	var teachers []entities.Teacher
	if result := repository.db.Limit(limit).Offset(offset).Find(&teachers); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get teachers : %s", result.Error))
	}

	return teachers, err
}

// AddTeacher 添加教师
func (repository *TeachersRepository) AddTeacher(teacher *entities.Teacher) (uint, error) {
	var err error
	if result := repository.db.Create(teacher); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add teacher %+v : %s", teacher, result.Error))
	}

	return teacher.ID, err
}

// UpdateTeacher 更新教师信息
func (repository *TeachersRepository) UpdateTeacher(teacher *entities.Teacher) error {
	var err error

	if result := repository.db.Save(&teacher); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update teacher %+v : %s", teacher, result.Error))
	}

	return err
}

// DeleteTeacher 删除教师
func (repository *TeachersRepository) DeleteTeacher(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.Teacher{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete teacher id %v : %s", id, result.Error))
	}

	return err
}

// TeacherExists 判断该教师是否存在
func (repository *TeachersRepository) TeacherExists(id uint) bool {
	var teacher entities.Teacher
	result := repository.db.First(&teacher, id)

	return result.RowsAffected >= 1
}

// Count 返回数量
func (repository *TeachersRepository) Count() int64 {
	var count int64
	repository.db.Model(&entities.Teacher{}).Count(&count)

	return count
}
