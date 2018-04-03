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

// GetAll 全部取得
func GetAll() []Sheet {
	var sheets = GetAllSheets()

	for _, v := range sheets {
		v.GetAllTodos()
	}

	return sheets
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
func (s Sheet) GetAllTodos() []Todo {
	db := NewDBConn()
	defer db.Close()

	var todos []Todo
	db.Where("sheet_id = ?", s.ID).Find(&todos)
	s.Todos = append(s.Todos, todos...)
	return todos
}

//Init 初期化
func Init() {

	db := NewDBConn()
	defer db.Close()

	// Migrate
	db.AutoMigrate(&Sheet{})
	db.AutoMigrate(&Todo{})
	/*

		// Create
		var todos = new([2]Todo)
		todos[0].ID = 1
		todos[0].Name = "頑張る？"
		todos[0].SheetID = 1
		todos[1].ID = 2
		todos[1].Name = "頑張れ・・・？"
		todos[1].SheetID = 1

		// 構造体を作成し
		var sts Sheet

		sts.ID = 1
		sts.Title = "テスト"

				// 中に入れる処理を書き
				db.First(&sts, 1)

				//Appendの使い方がOMOSIROI

			// 出力してみる
			db.Save(sts)
			db.Save(todos[0])
			db.Save(todos[1])

			var test Sheet
			var tds []Todo

			db.First(&test)
			db.Where("sheet_id = ?", test.ID).Find(&tds)
			test.Todos = append(test.Todos, tds...)
			//スライスしてやると追加できる（可変する場合はスライスを絶対にする）
			fmt.Println(test)
	*/

}
