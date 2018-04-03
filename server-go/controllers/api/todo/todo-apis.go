package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArakiTakaki/todo-lesson/server-go/models"
	"github.com/gin-gonic/gin"
)

// AppSet アプリケーションのルーティング
func AppSet(r *gin.RouterGroup) {
	r.POST("/create/todo", createTodo)
	r.GET("/sheets", getSheet)
	r.GET("/todos", getTodo)
}

func getSheet(c *gin.Context) {
	sheets := models.GetAllSheets()
	fmt.Println("GET ALL SHEETS")
	fmt.Println(sheets)
	c.JSON(http.StatusOK, sheets)
}

func getTodo(c *gin.Context) {
	id := c.Query("sheet_id")
	todo := models.GetAllTodos(id)
	c.JSON(http.StatusOK, todo)
}
func createTodo(c *gin.Context) {
	content := c.PostForm("data")
	sheet := c.PostForm("sheet")
	sheetID, _ := strconv.Atoi(sheet)
	todo := models.SetTodo(sheetID, content)
	c.JSON(http.StatusOK, todo)

}

/*
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
*/
