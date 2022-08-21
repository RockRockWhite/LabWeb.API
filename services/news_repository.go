package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

// NewNewsRepository 创建新的NewsRepository
func NewNewsRepository(autoMigrate bool) *NewsRepository {
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
		if err := db.AutoMigrate(&entities.News{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "News", err))
		}
	}

	repository := NewsRepository{db}
	return &repository
}

// GetNews 从id获得新闻
func (repository *NewsRepository) GetNews(id uint) (*entities.News, error) {
	var err error
	var news entities.News
	if result := repository.db.First(&news, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get news id %v : %s", id, result.Error))
	}

	return &news, err
}

// GetNewsList 获得新闻列表
func (repository *NewsRepository) GetNewsList(limit int, offset int) ([]entities.News, error) {
	var err error
	var news []entities.News
	if result := repository.db.Limit(limit).Offset(offset).Find(&news); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get news : %s", result.Error))
	}

	return news, err
}

// AddNews 添加新闻
func (repository *NewsRepository) AddNews(news *entities.News) (uint, error) {
	var err error
	if result := repository.db.Create(news); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add news %+v : %s", news, result.Error))
	}

	return news.ID, err
}

// UpdateNews 更新新闻
func (repository *NewsRepository) UpdateNews(news *entities.News) error {
	var err error

	if result := repository.db.Save(&news); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update news %+v : %s", news, result.Error))
	}

	return err
}

// DeleteNews 删除新闻
func (repository *NewsRepository) DeleteNews(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.News{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete news id %v : %s", id, result.Error))
	}

	return err
}

// NewsExists 判断该新闻是否存在
func (repository *NewsRepository) NewsExists(id uint) bool {
	var news entities.News
	result := repository.db.First(&news, id)

	return result.RowsAffected >= 1
}
