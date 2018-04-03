package models

import (
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
func GetAllTodos(sheetID int) []Todo {
	db := NewDBConn()
	defer db.Close()

	var todos []Todo
	db.Where("sheet_id = ?", sheetID).Find(&todos)
	return todos
}

//Init 初期化
func Init() {

	db := NewDBConn()
	defer db.Close()

	// Migrate
	db.AutoMigrate(&Sheet{})
	db.AutoMigrate(&Todo{})

	var sheet Sheet

	db.Create(&sheet)

	//スライスしてやると追加できる（可変する場合はスライスを絶対にする）

}
