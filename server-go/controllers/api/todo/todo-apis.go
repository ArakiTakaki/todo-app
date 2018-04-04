package todo

import (
	"net/http"

	"github.com/ArakiTakaki/todo-lesson/server-go/models"
	"github.com/gin-gonic/gin"
)

// AppSet アプリケーションのルーティング

// GetSheet nil
func GetSheet(c *gin.Context) {
	sheets := models.GetAllSheets()
	c.JSON(http.StatusOK, sheets)
}

// CreateSheet PostData : title str
func CreateSheet(c *gin.Context) {
	title := c.PostForm("title")
	sheet := models.CreateSheet(title)
	c.JSON(http.StatusOK, sheet)
}

// GetTodo /:sheet_id/
func GetTodo(c *gin.Context) {
	id := c.Query("sheet_id")
	todo := models.GetAllTodos(id)

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {

	var todo models.Todo

	content := c.PostForm("data")
	sheet := c.Param("sheet")
	todo = models.SetTodo(sheet, content)

	c.JSON(http.StatusOK, todo)
}

/*
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
*/
