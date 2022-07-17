package db

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type Config struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
}

var DB *gorm.DB

func init() {
	//	前提：create database bubble;

	var conf Config
	_, err := toml.DecodeFile("db/config.toml", &conf)
	if err != nil {
		log.Fatal(err)
	}

	//	connect to database
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//	create table: todos
	err = DB.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal(err)
	}
}

func AddToDo(todo Todo) {
	err := DB.Create(&todo).Error
	if err != nil {
		log.Fatal(err)
	}
}

//func FindToDo(id string, todo *Todo) {
//	err := DB.Where("id = ?", id).First(todo).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func FindAllToDo(todoList *[]Todo) {
	err := DB.Find(todoList).Error
	if err != nil {
		log.Fatal(err)
	}
}

func SaveToDo(id string, todo *Todo) {
	err := DB.Where("id = ?", id).First(todo).Update("status", !todo.Status).Error
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteToDo(id string, todo *Todo) {
	err := DB.Where("id = ?", id).First(todo).Delete(todo).Error
	if err != nil {
		log.Fatal(err)
	}
}
