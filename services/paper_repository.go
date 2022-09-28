package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type PapersRepository struct {
	db *gorm.DB
}

var _papersRepository *PapersRepository

func init() {
	db := getDB()
	if err := db.AutoMigrate(&entities.Paper{}); err != nil {
		utils.GetLogger().Fatal("Fatal migrate database %s : %s \n", "Paper", err)
	}

	_papersRepository = &PapersRepository{db}
}

func GetPapersRepository() *PapersRepository {
	return _papersRepository
}

// GetPaper  从id获得论文
func (repository *PapersRepository) GetPaper(id uint) (*entities.Paper, error) {
	var err error
	var paper entities.Paper
	if result := repository.db.First(&paper, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get paper id %v : %s", id, result.Error))
	}

	return &paper, err
}

// GetPapers 获得论文列表
func (repository *PapersRepository) GetPapers(limit int, offset int) ([]entities.Paper, error) {
	var err error
	var papers []entities.Paper
	if result := repository.db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&papers); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get papers : %s", result.Error))
	}

	return papers, err
}

// AddPaper 创建论文
func (repository *PapersRepository) AddPaper(paper *entities.Paper) (uint, error) {
	var err error
	if result := repository.db.Create(paper); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add paper %+v : %s", paper, result.Error))
	}

	return paper.ID, err
}

// UpdatePaper 更新论文
func (repository *PapersRepository) UpdatePaper(paper *entities.Paper) error {
	var err error

	if result := repository.db.Save(&paper); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update paper %+v : %s", paper, result.Error))
	}

	return err
}

// DeletePaper 删除论文
func (repository *PapersRepository) DeletePaper(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.Paper{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete paper id %v : %s", id, result.Error))
	}

	return err
}

// PaperExists 判断该id是否存在
func (repository *PapersRepository) PaperExists(id uint) bool {
	var paper entities.Paper
	result := repository.db.First(&paper, id)

	return result.RowsAffected >= 1
}

// Count 返回数量
func (repository *PapersRepository) Count() int64 {
	var count int64
	repository.db.Model(&entities.Paper{}).Count(&count)

	return count
}
