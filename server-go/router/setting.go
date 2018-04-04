package router

import (
	"github.com/ArakiTakaki/todo-lesson/server-go/controllers/api/todo"
	"github.com/gin-gonic/gin"
)

// 静的ファイルのディレクトリ記述場所
var static = map[string]string{
	"/css":   "./static/css",
	"/js":    "./static/js",
	"/image": "./static/image",
}

// setRouting
func setRouting(r *gin.Engine) {
	todoSet(r.Group("/api/todo"))
}
func todoSet(api *gin.RouterGroup) {
	api.GET("/sheets", todo.GetSheet)
	api.GET("/todos", todo.GetTodo)
	api.POST("/:sheet/create", todo.CreateTodo)
	api.POST("/create/newsheet", todo.CreateSheet)
}
