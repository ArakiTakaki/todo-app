package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRouter ルーターをゲットする
func GetRouter() *gin.Engine {

	var r = gin.Default()
	setStatic(r)
	setRouting(r)
	r.LoadHTMLGlob("views/*")
	r.NoRoute(func(c *gin.Context) {
		fmt.Println("indexに飛ばします")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	return r
}

// SetStatic スタティックファイルの設定
func setStatic(r *gin.Engine) {
	for k, v := range static {
		r.Static(k, v)
	}
}
