package models

import "time"

// Model init
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  time.Time  `json:"update_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

//Sheet シート
type Sheet struct {
	Model
	Title string `json:"sheet_title"`
	Todos []Todo `json:"todos"`
}

//Todo 記載用
type Todo struct {
	Model
	SheetID uint   `gorm:"index" json:"sheet_id"`
	Name    string `json:"todo_name"`
	More    string `json:"todo_more"`
	Check   bool   `json:"todo_check"`
}

/*

# 不要な気がする

// GetSheet そのまま return *Sheet
func GetSheet() Sheet {
	var sheet Sheet
	return sheet
}

// GetSheets そのまま return *Sheet
func GetSheets() []Sheet {
	var sheets []Sheet
	return sheets
}

// GetTodo そのまま return *Todo
func GetTodo() Todo {
	return Todo{}
}

// GetTodos そのまま return *Todo
func GetTodos(n int) []Todo {
	var todos []Todo
	return todos
}
*/
