package main

import (
	"bubble/db"
	"bubble/routers"
)

func main() {
	db.Init()

	r := routers.SetUpRouter()
	err := r.Run()
	if err != nil {
		return
	}
}
