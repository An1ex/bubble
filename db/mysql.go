package db

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
}

var DB *gorm.DB

func connect() {
	var conf Config
	_, err := toml.DecodeFile("config/config.toml", &conf)
	if err != nil {
		log.Fatal(err)
	}

	//	connect to database
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Init() {
	//	前提：create database bubble;
	connect()
	//	create table todos
	err := DB.AutoMigrate(&Todo{})
	if err != nil {
		panic(err)
	}
}
