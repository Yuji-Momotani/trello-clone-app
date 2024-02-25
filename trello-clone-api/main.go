package main

import (
	"trello-colen-api/controller"
	"trello-colen-api/db"
	"trello-colen-api/repository"
	"trello-colen-api/router"
	"trello-colen-api/usecase"
	"trello-colen-api/validator"
)

func main() {
	connectDB := db.NewDB()

	// user
	userValidation := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(connectDB)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidation)
	userController := controller.NewUserController(userUsecase)

	// task
	taskRepository := repository.NewTaskRepository(connectDB)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	// router
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
