package models

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	//SQLite3 ドライバ
	_ "github.com/mattn/go-sqlite3"
)

// NewDBConn  データベースへのアクセスオブジェクト
func NewDBConn() *gorm.DB {
	const CONNECT = "todo.sqlite3"
	db, err := gorm.Open("sqlite3", CONNECT)
	if err != nil {
		panic(err)
	}
	return db
}

// GetAllSheets ALL
func GetAllSheets() []Sheet {
	db := NewDBConn()
	defer db.Close()

	var sheets []Sheet

	db.Where("delete_at IS NULL").Find(&sheets)

	return sheets
}

// GetAllTodos ALL
func GetAllTodos(sheetID string) []Todo {
	db := NewDBConn()
	defer db.Close()

	var todos []Todo
	db.Where("sheet_id = ?", sheetID).Find(&todos)
	fmt.Println(todos)
	return todos
}

// SetTodo セットする
func SetTodo(sheetID string, content string) Todo {
	db := NewDBConn()
	defer db.Close()

	var todo Todo
	work, _ := strconv.Atoi(sheetID)
	i := uint(work)

	todo.SheetID = i
	todo.Name = content

	db.Save(&todo)
	return (todo)
}

//Init 初期化
func Init() {

	db := NewDBConn()
	defer db.Close()

	// Migrate
	db.AutoMigrate(&Sheet{})
	db.AutoMigrate(&Todo{})

	/*
		var sheet Sheet
		sheet.Title = "テスト"
		db.Create(&sheet)
	*/

	//スライスしてやると追加できる（可変する場合はスライスを絶対にする）

}
