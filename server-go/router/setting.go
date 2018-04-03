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
	todo.AppSet(r.Group("/api/todo"))
}
