package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repositoty"
	"go-rest-api/router"
	"go-rest-api/usecase"
)

func main() {
	// DB
	db := db.NewDB()

	// ユーザ関連 ---------------------------------
	// レポジトリ
	userRepository := repositoty.NewUserRepository(db)
	// ユースケース
	userUsecase := usecase.NewUserUsecase(userRepository)
	// コントローラー
	userController := controller.NewUserController(userUsecase)

	// タスク関連 ---------------------------------
	taskRepository := repositoty.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	// ルーター
	e := router.NewRouter(userController, taskController)

	// サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}