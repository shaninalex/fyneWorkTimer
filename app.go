package main

import (
	"database/sql"
)

type App struct {
	DB *sql.DB
}

func (app *App) Init(dbpath string) {
	db, err := InitDatabase(dbpath)
	if err != nil {
		panic(err)
	}
	app.DB = db
}
