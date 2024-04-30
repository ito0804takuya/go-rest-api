// migrate実行時には、main関数を実行したい
// そのためにはmainパッケージに所属させないといけない
package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()

	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)

	// AutoMigrate() を実行すると、その構造体に対応するテーブルが作成される
	// NOTE: AutoMigrateは新規作成のみ実行する。つまり、スキーマの変更などは、別で処理しないといけない
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
