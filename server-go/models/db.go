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

func CreateSheet(title string) Sheet {
	db := NewDBConn()
	defer db.Close()

	var sheet Sheet
	sheet.Title = title

	db.Create(&sheet)

	return sheet
}

// GetAllTodos ALL
func GetAllTodos(sheetID string) []Todo {
	db := NewDBConn()
	defer db.Close()

	var todos []Todo
	db.Where("sheet_id = ?", sheetID).Find(&todos)

	return todos
}

// SetTodo セットする
func SetTodo(strSheetID string, content string) Todo {
	intSheetID, err := strconv.Atoi(strSheetID)
	uintSheetID := uint(intSheetID)
	var todo Todo
	if err != nil {
		fmt.Println("不正な値が検出されました setTodo")
		return todo
	}

	db := NewDBConn()
	defer db.Close()

	todo.SheetID = uintSheetID
	todo.Name = content
	todo.Check = false

	db.Create(&todo)

	return todo
}

//Init 初期化
func Init() {

	db := NewDBConn()
	defer db.Close()

	// Migrate
	db.AutoMigrate(&Sheet{})
	db.AutoMigrate(&Todo{})

}
