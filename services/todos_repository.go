package services

import (
	"fmt"
	"github.com/RockRockWhite/LabWeb.API/entities"
	"github.com/RockRockWhite/LabWeb.API/utils"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type TodosRepository struct {
	db *gorm.DB
}

var _todosRepository *TodosRepository

func init() {
	db := getDB()
	if err := db.AutoMigrate(&entities.Todo{}); err != nil {
		utils.GetLogger().Fatal("Fatal migrate database %s : %s \n", "Todos", err)
	}

	_todosRepository = &TodosRepository{db}
}

func GetTodoRepository() *TodosRepository {
	return _todosRepository
}

// GetTodo 从id获得todo
func (repository *TodosRepository) GetTodo(id uint) (*entities.Todo, error) {
	var err error
	var todos entities.Todo
	if result := repository.db.First(&todos, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get todos id %v : %s", id, result.Error))
	}

	return &todos, err
}

// GetTodosList 获得todo列表
func (repository *TodosRepository) GetTodosList(limit int, offset int) ([]entities.Todo, error) {
	var err error
	var todos []entities.Todo
	if result := repository.db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&todos); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to get todos : %s", result.Error))
	}

	return todos, err
}

// AddTodo 添加todo
func (repository *TodosRepository) AddTodo(todo *entities.Todo) (uint, error) {
	var err error
	if result := repository.db.Create(todo); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to add todo %+v : %s", todo, result.Error))
	}

	return todo.ID, err
}

// UpdateTodo 更新todo
func (repository *TodosRepository) UpdateTodo(todo *entities.Todo) error {
	var err error

	if result := repository.db.Save(&todo); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to update todo %+v : %s", todo, result.Error))
	}

	return err
}

// DeleteTodo 删除todo
func (repository *TodosRepository) DeleteTodo(id uint) error {
	var err error

	if result := repository.db.Delete(&entities.Todo{}, id); result.Error != nil {
		err = multierror.Append(err, fmt.Errorf("failed to delete todo id %v : %s", id, result.Error))
	}

	return err
}

// TodoExists 判断该todo是否存在
func (repository *TodosRepository) TodoExists(id uint) bool {
	var todo entities.Todo
	result := repository.db.First(&todo, id)

	return result.RowsAffected >= 1
}

// Count 返回数量
func (repository *TodosRepository) Count() int64 {
	var count int64
	repository.db.Model(&entities.Todo{}).Count(&count)

	return count
}
