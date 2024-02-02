package main

import (
	"github.com/yablus/les28/pkg/storage"
	"github.com/yablus/les28/pkg/student"
)

func main() {
	//repository := &student.DataTest{}
	//repository := storage.NewStorage()
	app := &student.App{storage.NewStorage()}
	app.Run()
}
