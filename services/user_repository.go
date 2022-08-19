package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建新用户Repository
func NewUserRepository(autoMigrate bool) *UserRepository {
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

	// 完成User迁移
	if autoMigrate {
		if err := db.AutoMigrate(&entities.User{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "User", err))
		}
	}

	repository := UserRepository{db}
	return &repository
}

// GetUserByName 通过用户名获得信息
func (repository *UserRepository) GetUserByName(username string) (*entities.User, error) {
	var err error
	var user entities.User

	if result := repository.db.Where(&entities.User{Username: username}).First(&user); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get user: %v, error: %s", username, err.Error()))
	}

	return &user, err
}

// AddUser 添加用户
func (repository *UserRepository) AddUser(user *entities.User) (uint, error) {
	var err error
	if result := repository.db.Create(user); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add user %+v : %s", user, result.Error))
	}

	return user.ID, err
}

// UpdateUser 更新用户信息
func (repository *UserRepository) UpdateUser(user *entities.User) error {
	var err error
	if result := repository.db.Save(user); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update user %+v : %s", user, result.Error))
	}

	return err
}

// DeleteUserByName 通过用户名删除用户
func (repository *UserRepository) DeleteUserByName(username string) error {
	var err error
	if result := repository.db.Where("username = ?", username).Delete(&entities.User{}); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete user %+v : %s", username, result.Error))
	}

	return err
}

// UserIdExists 判断用户id是否存在
func (repository *UserRepository) UserIdExists(id uint) bool {
	var user entities.User
	result := repository.db.First(&user, id)

	return result.RowsAffected >= 1
}

// UsernameExists 判断用户昵称是否存在
func (repository *UserRepository) UsernameExists(username string) bool {
	var user entities.User
	result := repository.db.Where(&entities.User{Username: username}).First(&user)

	return result.RowsAffected >= 1
}
