package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	//	静态文件路径
	r.Static("/static", "./static")

	//	模板文件路径
	r.LoadHTMLGlob("./templates/*")

	//	router index
	r.GET("/", controller.Index)

	// router group
	v1Group := r.Group("v1")

	//	查看所有代办事项
	v1Group.GET("/todo", controller.GetToDoList)

	//	添加代办事项
	v1Group.POST("/todo", controller.CreateToDo)

	//	修改代办事项
	v1Group.PUT("/todo/:id", controller.UpdateToDo)

	//	删除代办事项
	v1Group.DELETE("/todo/:id", controller.DeleteToDo)

	return r
}
