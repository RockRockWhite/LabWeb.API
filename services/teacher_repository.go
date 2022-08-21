package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	db *gorm.DB
}

// NewTeacherRepository 创建新的TeacherRepository
func NewTeacherRepository(autoMigrate bool) *TeacherRepository {
	Host := viper.GetString("DataBase.Host")
	Port := viper.GetString("DataBase.Port")
	Username := viper.GetString("DataBase.Username")
	Password := viper.GetString("DataBase.Password")
	DBName := viper.GetString("DataBase.DBName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error open database:%s %s \n", dsn, err))
	}

	// 完成Article迁移
	if autoMigrate {
		if err := db.AutoMigrate(&entities.Teacher{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Teacher", err))
		}
	}

	repository := TeacherRepository{db}
	return &repository
}

// GetTeacher 从id获得教师
func (repository *TeacherRepository) GetTeacher(id uint) (*entities.Teacher, error) {
	var err error
	var teacher entities.Teacher
	if result := repository.db.First(&teacher, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get teacher id %v : %s", id, result.Error))
	}

	return &teacher, err
}

// GetTeachers 获得教师列表
func (repository *TeacherRepository) GetTeachers(limit int, offset int) ([]entities.Teacher, error) {
	var err error
	var teachers []entities.Teacher
	if result := repository.db.Limit(limit).Offset(offset).Find(&teachers); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get teachers : %s", result.Error))
	}

	return teachers, err
}

// AddTeacher 添加教师
func (repository *TeacherRepository) AddTeacher(teacher *entities.Teacher) (uint, error) {
	var err error
	if result := repository.db.Create(teacher); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add teacher %+v : %s", teacher, result.Error))
	}

	return teacher.ID, err
}

// UpdateTeacher 更新教师信息
func (repository *TeacherRepository) UpdateTeacher(teacher *entities.Teacher) error {
	var err error

	if result := repository.db.Save(&teacher); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update teacher %+v : %s", teacher, result.Error))
	}

	return err
}

// DeleteTeacher 删除教师
func (repository *TeacherRepository) DeleteTeacher(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.Teacher{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete teacher id %v : %s", id, result.Error))
	}

	return err
}

// TeacherExists 判断该教师是否存在
func (repository *TeacherRepository) TeacherExists(id uint) bool {
	var teacher entities.Teacher
	result := repository.db.First(&teacher, id)

	return result.RowsAffected >= 1
}
