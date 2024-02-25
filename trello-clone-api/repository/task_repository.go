package repository

import (
	"errors"
	"trello-colen-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// インターフェース
type ITaskRepository interface {
	GetTaskCardsByUser(task_cards *[]model.TaskCard, user_id uint) error
	GetTasksByTaskCard(task *[]model.Task, task_card_id uint) error
	RegistTaskCard(task_card *model.TaskCard) error
	UpdateTaskCard(task_card *model.TaskCard, id uint, user_id uint) error
	DeleteTaskCard(id uint, user_id uint) error
	RegistTask(task *model.Task) error
	UpdateTask(task *model.Task, id uint) error
	DeleteTask(id uint) error
}

// 構造体
type taskRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetTaskCardsByUser(task_cards *[]model.TaskCard, user_id uint) error {
	if err := tr.db.Where("user_id = ?", user_id).Order("sort_no").Find(task_cards).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTasksByTaskCard(task *[]model.Task, task_card_id uint) error {
	if err := tr.db.Where("task_card_id = ?", task_card_id).Order("sort_no").Find(task).Error; err != nil {
		return err
	}
	return nil
}

// ******************************************
// Task_Cardの処理
// ******************************************
func (tr *taskRepository) RegistTaskCard(task_card *model.TaskCard) error {
	if err := tr.db.Create(task_card).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTaskCard(task_card *model.TaskCard, id uint, user_id uint) error {
	result := tr.db.Model(task_card).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", id, user_id).Updates(map[string]interface{}{
		"title":   task_card.Title,
		"sort_no": task_card.SortNo,
	})

	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows < 1 {
		return errors.New("record not found")
	}

	return nil
}

func (tr *taskRepository) DeleteTaskCard(id uint, user_id uint) error {
	taskCard := new(model.TaskCard)
	result := tr.db.Where("id = ? AND user_id = ?", id, user_id).Delete(taskCard)
	if err := result.Error; err != nil {
		return err
	}

	if rows := result.RowsAffected; rows < 1 {
		return errors.New("record not found")
	}
	return nil
}

// ******************************************
// Taskの処理
// ******************************************
func (tr *taskRepository) RegistTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, id uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ?", id).Updates(map[string]interface{}{
		"content": task.Content,
		"sort_no": task.SortNo,
	})

	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows < 1 {
		return errors.New("record not found")
	}

	return nil
}

func (tr *taskRepository) DeleteTask(id uint) error {
	task := new(model.Task)
	result := tr.db.Where("id = ?", id).Delete(task)
	if err := result.Error; err != nil {
		return err
	}

	if rows := result.RowsAffected; rows < 1 {
		return errors.New("record not found")
	}
	return nil
}
