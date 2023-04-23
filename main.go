package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var app = WorkTimer{}
	app.Init("timer.db")
	app.Run()
}
