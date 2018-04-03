package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArakiTakaki/todo-lesson/server-go/models"
	"github.com/gin-gonic/gin"
)

// AppSet アプリケーションのルーティング

func GetSheet(c *gin.Context) {
	sheets := models.GetAllSheets()
	fmt.Println("GET ALL SHEETS")
	fmt.Println(sheets)
	c.JSON(http.StatusOK, sheets)
}

func GetTodo(c *gin.Context) {
	id := c.Query("sheet_id")
	todo := models.GetAllTodos(id)
	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {

	var todo models.Todo

	content := c.PostForm("data")
	sheet := c.Param("sheet")
	sheetID, err := strconv.Atoi(sheet)

	if err == nil {
		todo = models.SetTodo(sheetID, content)
	} else {
		fmt.Println("todo-api create-todo :不正な値が検出されました" + sheet)
	}

	c.JSON(http.StatusOK, todo)
}

/*
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
*/
