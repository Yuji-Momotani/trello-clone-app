package usecase

import (
	"trello-colen-api/model"
	"trello-colen-api/repository"
)

// インターフェース
type ITaskUsecase interface {
	GetTaskCardsAndTasks(user_id uint) ([]model.TaskCardAndTasksResponse, error)
	RegistTaskCard(task_card model.TaskCard, user_id uint) (model.TaskCardResponse, error)
	UpdateTaskCard(task_card model.TaskCard, user_id uint, id uint) error
	DeleteTaskCard(user_id uint, id uint) error
	RegistTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, id uint) error
	DeleteTask(id uint) error
}

// 構造体
type taskUsecase struct {
	tr repository.ITaskRepository
}

// コンストラクタ
func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func ConvertTaskToTaskResponse(tasks []model.Task) []model.TaskResponse {
	var tasksResponse []model.TaskResponse
	for _, task := range tasks {
		taskResponse := model.TaskResponse{
			ID:         task.ID,
			Contetn:    task.Content,
			TaskCardID: task.TaskCardID,
			SortNo:     task.SortNo,
		}
		tasksResponse = append(tasksResponse, taskResponse)
	}
	return tasksResponse
}

// 時間の関係上、一旦validatorは省略。時間があれば実装する。（2/12）

func (tu *taskUsecase) GetTaskCardsAndTasks(user_id uint) ([]model.TaskCardAndTasksResponse, error) {
	taskCards := []model.TaskCard{}
	if err := tu.tr.GetTaskCardsByUser(&taskCards, user_id); err != nil {
		return []model.TaskCardAndTasksResponse{}, err
	}

	taskCardsTasksRes := []model.TaskCardAndTasksResponse{}
	for _, card := range taskCards {
		tasks := []model.Task{}

		if err := tu.tr.GetTasksByTaskCard(&tasks, card.ID); err != nil {
			return []model.TaskCardAndTasksResponse{}, err
		}
		taskResponse := ConvertTaskToTaskResponse(tasks)
		taskCardsTasksRes = append(taskCardsTasksRes, model.TaskCardAndTasksResponse{
			ID:     card.ID,
			Titlle: card.Title,
			SortNo: card.SortNo,
			Tasks:  taskResponse,
		})
	}
	return taskCardsTasksRes, nil
}

// ******************************************
// Task_Cardの処理
// ******************************************
func (tu *taskUsecase) RegistTaskCard(task_card model.TaskCard, user_id uint) (model.TaskCardResponse, error) {
	registTaskCard := model.TaskCard{
		Title:  task_card.Title,
		UserID: user_id,
		SortNo: task_card.SortNo,
	}

	if err := tu.tr.RegistTaskCard(&registTaskCard); err != nil {
		return model.TaskCardResponse{}, err
	}

	taskCardResponse := model.TaskCardResponse{
		ID:     registTaskCard.ID,
		Titlle: registTaskCard.Title,
		SortNo: registTaskCard.SortNo,
	}
	return taskCardResponse, nil
}

func (tu *taskUsecase) UpdateTaskCard(task_card model.TaskCard, user_id uint, id uint) error {
	if err := tu.tr.UpdateTaskCard(&task_card, id, user_id); err != nil {
		return err
	}
	return nil
}

func (tu *taskUsecase) DeleteTaskCard(user_id uint, id uint) error {
	if err := tu.tr.DeleteTaskCard(id, user_id); err != nil {
		return err
	}

	return nil
}

// ******************************************
// Taskの処理
// ******************************************
func (tu *taskUsecase) RegistTask(task model.Task) (model.TaskResponse, error) {
	registTask := model.Task{
		Content:    task.Content,
		TaskCardID: task.TaskCardID,
		SortNo:     task.SortNo,
	}

	if err := tu.tr.RegistTask(&registTask); err != nil {
		return model.TaskResponse{}, err
	}

	taskResponse := model.TaskResponse{
		ID:         registTask.ID,
		Contetn:    registTask.Content,
		TaskCardID: registTask.TaskCardID,
		SortNo:     registTask.SortNo,
	}
	return taskResponse, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, id uint) error {
	if err := tu.tr.UpdateTask(&task, id); err != nil {
		return err
	}
	return nil
}

func (tu *taskUsecase) DeleteTask(id uint) error {
	if err := tu.tr.DeleteTask(id); err != nil {
		return err
	}

	return nil
}
