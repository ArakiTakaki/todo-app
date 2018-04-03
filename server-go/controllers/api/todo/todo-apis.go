package todo

import (
	"fmt"
	"net/http"

	"github.com/ArakiTakaki/todo-lesson/server-go/models"
	"github.com/gin-gonic/gin"
)

func AppSet(r *gin.RouterGroup) {
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
	fmt.Println(id)
	c.JSON(http.StatusOK, nil)
}

/*
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
*/
