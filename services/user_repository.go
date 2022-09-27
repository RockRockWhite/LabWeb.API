package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

var _usersRepository *UsersRepository

func init() {
	db := getDB()
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		utils.GetLogger().Fatal("Fatal migrate database %s : %s \n", "User", err)
	}

	_usersRepository = &UsersRepository{db}
}

func GetUsersRepository() *UsersRepository {
	return _usersRepository
}

// GetUserByName 通过用户名获得信息
func (repository *UsersRepository) GetUserByName(username string) (*entities.User, error) {
	var err error
	var user entities.User

	if result := repository.db.Where(&entities.User{Username: username}).First(&user); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get user: %v, error: %s", username, err.Error()))
	}

	return &user, err
}

// AddUser 添加用户
func (repository *UsersRepository) AddUser(user *entities.User) (uint, error) {
	var err error
	if result := repository.db.Create(user); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add user %+v : %s", user, result.Error))
	}

	return user.ID, err
}

// UpdateUser 更新用户信息
func (repository *UsersRepository) UpdateUser(user *entities.User) error {
	var err error
	if result := repository.db.Save(user); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update user %+v : %s", user, result.Error))
	}

	return err
}

// DeleteUserByName 通过用户名删除用户
func (repository *UsersRepository) DeleteUserByName(username string) error {
	var err error
	if result := repository.db.Where("username = ?", username).Delete(&entities.User{}); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete user %+v : %s", username, result.Error))
	}

	return err
}

// UserIdExists 判断用户id是否存在
func (repository *UsersRepository) UserIdExists(id uint) bool {
	var user entities.User
	result := repository.db.First(&user, id)

	return result.RowsAffected >= 1
}

// UsernameExists 判断用户昵称是否存在
func (repository *UsersRepository) UsernameExists(username string) bool {
	var user entities.User
	result := repository.db.Where(&entities.User{Username: username}).First(&user)

	return result.RowsAffected >= 1
}
