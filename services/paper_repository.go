package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PaperRepository struct {
	db *gorm.DB
}

// NewPaperRepository 创建新的PaperRepository
func NewPaperRepository(autoMigrate bool) *PaperRepository {
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
		if err := db.AutoMigrate(&entities.Paper{}); err != nil {
			panic(fmt.Errorf("Fatal migrate database %s : %s \n", "Article", err))
		}
	}

	repository := PaperRepository{db}
	return &repository
}

// GetPaper  从id获得论文
func (repository *PaperRepository) GetPaper(id uint) (*entities.Paper, error) {
	var err error
	var paper entities.Paper
	if result := repository.db.First(&paper, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get paper id %v : %s", id, result.Error))
	}

	return &paper, err
}

// GetPapers 获得论文列表
func (repository *PaperRepository) GetPapers(limit int, offset int) ([]entities.Paper, error) {
	var err error
	var papers []entities.Paper
	if result := repository.db.Limit(limit).Offset(offset).Find(&papers); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get papers : %s", result.Error))
	}

	return papers, err
}

// AddPaper 创建论文
func (repository *PaperRepository) AddPaper(paper *entities.Paper) (uint, error) {
	var err error
	if result := repository.db.Create(paper); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add paper %+v : %s", paper, result.Error))
	}

	return paper.ID, err
}

// UpdatePaper 更新论文
func (repository *PaperRepository) UpdatePaper(paper *entities.Paper) error {
	var err error

	if result := repository.db.Save(&paper); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update paper %+v : %s", paper, result.Error))
	}

	return err
}

// DeletePaper 删除论文
func (repository *PaperRepository) DeletePaper(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.Paper{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete paper id %v : %s", id, result.Error))
	}

	return err
}

// PaperExists 判断该id是否存在
func (repository *PaperRepository) PaperExists(id uint) bool {
	var paper entities.Paper
	result := repository.db.First(&paper, id)

	return result.RowsAffected >= 1
}
