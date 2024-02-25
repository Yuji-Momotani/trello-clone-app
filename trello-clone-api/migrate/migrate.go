package main

import (
	"fmt"
	"trello-colen-api/db"
	"trello-colen-api/model"
	// 各環境に合わせてdb、modelをimportする
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// 例：userとtaskテーブルを作成したい場合
	dbConn.AutoMigrate(&model.User{}, &model.TaskCard{}, &model.Task{}) //作成したいモデルのstructを0値で引数に渡す
}
