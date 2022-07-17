package main

import (
	"bubble/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//	静态文件路径
	r.Static("/static", "./static")
	//	模板文件路径
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// router group
	v1Group := r.Group("v1")

	//	添加代办事项
	v1Group.POST("/todo", func(c *gin.Context) {
		var todo db.Todo
		//	1.从请求中取数据
		err := c.BindJSON(&todo)
		if err != nil {
			return
		}
		//	2.存入数据库
		db.AddToDo(todo)
		//	3.返回响应
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"msg":  "add ToDo success",
		//	"data": todo,
		//})
	})

	//	查看所有代办事项
	v1Group.GET("/todo", func(c *gin.Context) {
		var todoList []db.Todo
		db.FindAllToDo(&todoList)
		c.JSON(http.StatusOK, todoList)
		//c.JSON(http.StatusOK, gin.H{
		//	"msg":  "find ToDoList success",
		//	"data": todoList,
		//})
	})

	//	修改代办事项
	v1Group.PUT("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": "invalid id",
			})
		}

		var todo db.Todo
		db.SaveToDo(id, &todo)
	})

	//	删除代办事项
	v1Group.DELETE("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": "invalid id",
			})
		}

		var todo db.Todo
		db.DeleteToDo(id, &todo)
	})

	err := r.Run()
	if err != nil {
		return
	}
}
