package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type ResourcesRepository struct {
	db *gorm.DB
}

var _resourcesRepository *ResourcesRepository

func init() {
	db := getDB()
	if err := db.AutoMigrate(&entities.Resource{}); err != nil {
		utils.GetLogger().Fatal("Fatal migrate database %s : %s \n", "Resource", err)
	}

	_resourcesRepository = &ResourcesRepository{db}
}

func GetResourcesRepository() *ResourcesRepository {
	return _resourcesRepository
}

// GetResource  从id获得资源
func (repository *ResourcesRepository) GetResource(id uint) (*entities.Resource, error) {
	var err error
	var resource entities.Resource
	if result := repository.db.First(&resource, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get resource id %v : %s", id, result.Error))
	}

	return &resource, err
}

// GetResources 获得资源列表
func (repository *ResourcesRepository) GetResources(limit int, offset int, filter string) ([]entities.Resource, error) {
	var err error
	var resources []entities.Resource

	db := repository.db
	if filter != "" {
		db = db.Where("state = ?", filter)
	}
	if result := db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&resources); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get resources : %s", result.Error))
	}

	return resources, err
}

// AddResource 创建资源
func (repository *ResourcesRepository) AddResource(resource *entities.Resource) (uint, error) {
	var err error
	if result := repository.db.Create(resource); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add resource %+v : %s", resource, result.Error))
	}

	return resource.ID, err
}

// UpdateResource 更新资源
func (repository *ResourcesRepository) UpdateResource(resource *entities.Resource) error {
	var err error

	if result := repository.db.Save(&resource); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update resource %+v : %s", resource, result.Error))
	}

	return err
}

// DeleteResource 删除资源
func (repository *ResourcesRepository) DeleteResource(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.Resource{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete resource id %v : %s", id, result.Error))
	}

	return err
}

// ResourceExists 判断该id是否存在
func (repository *ResourcesRepository) ResourceExists(id uint) bool {
	var resource entities.Resource
	result := repository.db.First(&resource, id)

	return result.RowsAffected >= 1
}

// Count 返回数量
func (repository *ResourcesRepository) Count() int64 {
	var count int64
	repository.db.Model(&entities.Resource{}).Count(&count)

	return count
}
