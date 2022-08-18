package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
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

// GetUser 获得用户信息
func (repository *UserRepository) GetUser(id uint) *entities.User {
	if !repository.UserExists(id) {
		panic(fmt.Errorf("user id %v not exists", id))
	}

	var user entities.User
	if result := repository.db.First(&user, id); result.Error != nil {
		panic(fmt.Errorf("failed to get user id: %v", id))
	}

	return &user
}

// AddUser 添加用户
func (repository *UserRepository) AddUser(user *entities.User) uint {
	if result := repository.db.Create(user); result.Error != nil {
		panic(fmt.Errorf("failed to add user %+v : %s", user, result.Error))
	}

	return user.ID
}

// UpdateUser 更新用户信息
func (repository *UserRepository) UpdateUser(user *entities.User) {
	if result := repository.db.Save(user); result.Error != nil {
		panic(fmt.Errorf("failed to update user %+v : %s", user, result.Error))
	}
}

// DeleteUser 删除用户
func (repository *UserRepository) DeleteUser(id uint) {
	if !repository.UserExists(id) {
		panic(fmt.Errorf("user id %v not exists", id))
	}

	if result := repository.db.Delete(&entities.User{}, id); result.Error != nil {
		panic(fmt.Errorf("failed to delete user id %v : %s", id, result.Error))
	}
}

// UserExists 判断用户是否存在
func (repository *UserRepository) UserExists(id uint) bool {
	var user entities.User
	result := repository.db.First(&user, id)

	return result.RowsAffected >= 1
}

// GetUserByNickName 通过昵称获得用户信息
func (repository *UserRepository) GetUserByNickName(username string) *entities.User {

	if !repository.UsernameExists(username) {
		panic(fmt.Errorf("user nick name %v not exists", username))
	}

	var user entities.User
	if result := repository.db.Where(&entities.User{Username: username}).First(&user); result.Error != nil {
		panic(fmt.Errorf("failed to get user username: %v", username))
	}

	return &user
}

// UsernameExists 判断用户昵称是否存在
func (repository *UserRepository) UsernameExists(username string) bool {
	var user entities.User
	result := repository.db.Where(&entities.User{Username: username}).First(&user)

	return result.RowsAffected >= 1
}
