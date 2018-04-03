package main

import "github.com/ArakiTakaki/todo-lesson/server-go/router"

func main() {
	r := router.GetRouter()
	//models.Init()
	r.Run(":8000")
}
