package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PublicationRepository struct {
	db *gorm.DB
}

// NewPublicationRepository 创建新的PublicationRepository
func NewPublicationRepository(autoMigrate bool) *PublicationRepository {
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
		if err := db.AutoMigrate(&entities.Publication{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Article", err))
		}
	}

	repository := PublicationRepository{db}
	return &repository
}

// GetPublication  从id获得论文
func (repository *PublicationRepository) GetPublication(id uint) (*entities.Publication, error) {
	var err error
	var publication entities.Publication
	if result := repository.db.First(&publication, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get publication id %v : %s", id, result.Error))
	}

	return &publication, err
}

// GetPublications 获得论文列表
func (repository *PublicationRepository) GetPublications(limit int, offset int) ([]entities.Publication, error) {
	var err error
	var publications []entities.Publication
	if result := repository.db.Limit(limit).Offset(offset).Find(&publications); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get publications : %s", result.Error))
	}

	return publications, err
}

// AddPublication 创建论文
func (repository *PublicationRepository) AddPublication(publication *entities.Publication) (uint, error) {
	var err error
	if result := repository.db.Create(publication); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add publication %+v : %s", publication, result.Error))
	}

	return publication.ID, err
}

// UpdatePublication 更新论文
func (repository *PublicationRepository) UpdatePublication(publication *entities.Publication) error {
	var err error

	if result := repository.db.Save(&publication); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update publication %+v : %s", publication, result.Error))
	}

	return err
}

// DeletePublication 删除论文
func (repository *PublicationRepository) DeletePublication(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.Publication{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete publication id %v : %s", id, result.Error))
	}

	return err
}

// PublicationExists 判断该id是否存在
func (repository *PublicationRepository) PublicationExists(id uint) bool {
	var publication entities.Publication
	result := repository.db.First(&publication, id)

	return result.RowsAffected >= 1
}
