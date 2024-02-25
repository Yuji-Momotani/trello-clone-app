package controller

import (
	"net/http"
	"strconv"
	"trello-colen-api/model"
	"trello-colen-api/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// インターフェース
type ITaskController interface {
	GetTaskCardsAndTasks(c echo.Context) error
	RegistTaskCard(c echo.Context) error
	UpdateTaskCard(c echo.Context) error
	DeleteTaskCard(c echo.Context) error
	RegistTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

// 構造体
type taskController struct {
	tu usecase.ITaskUsecase
}

// コンストラクタ
func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func getUserIdFromJWT(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(float64)

	return uint(user_id)
}

func (tc *taskController) GetTaskCardsAndTasks(c echo.Context) error {
	user_id := getUserIdFromJWT(c)
	taskCardsAndTasks, err := tc.tu.GetTaskCardsAndTasks(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskCardsAndTasks)
}

func (tc *taskController) RegistTaskCard(c echo.Context) error {
	user_id := getUserIdFromJWT(c)
	taskCard := model.TaskCard{}

	if err := c.Bind(&taskCard); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resRegistTaskCard, err := tc.tu.RegistTaskCard(taskCard, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resRegistTaskCard)
}

func (tc *taskController) UpdateTaskCard(c echo.Context) error {
	user_id := getUserIdFromJWT(c)
	s_task_card_id := c.Param("taskCardId")
	task_card_id, _ := strconv.Atoi(s_task_card_id)
	taskCard := model.TaskCard{}

	if err := c.Bind(&taskCard); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := tc.tu.UpdateTaskCard(taskCard, user_id, uint(task_card_id)); err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (tc *taskController) DeleteTaskCard(c echo.Context) error {
	user_id := getUserIdFromJWT(c)
	s_task_card_id := c.Param("taskCardId")
	task_card_id, _ := strconv.Atoi(s_task_card_id)

	if err := tc.tu.DeleteTaskCard(user_id, uint(task_card_id)); err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (tc *taskController) RegistTask(c echo.Context) error {
	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resRegistTask, err := tc.tu.RegistTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resRegistTask)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	s_task_id := c.Param("taskId")
	task_id, _ := strconv.Atoi(s_task_id)
	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := tc.tu.UpdateTask(task, uint(task_id)); err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	s_task_id := c.Param("taskId")
	task_id, _ := strconv.Atoi(s_task_id)

	if err := tc.tu.DeleteTask(uint(task_id)); err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
