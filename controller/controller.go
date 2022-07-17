package controller

import (
	"bubble/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateToDo(c *gin.Context) {
	var todo db.Todo
	//	1.从请求中取数据
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	//	2.存入数据库
	err = db.Create(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	//	3.返回响应
	c.JSON(http.StatusOK, todo)
}

func GetToDoList(c *gin.Context) {
	todoList, err := db.ReadAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateToDo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "invalid id",
		})
	}
	err := db.Update(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteToDo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "invalid id",
		})
	}
	err := db.Delete(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
