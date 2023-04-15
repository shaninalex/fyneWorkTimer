package main

import (
	"database/sql"
	"log"
)

type App struct {
	DB *sql.DB
	ui *Ui
}

func (app *App) Init(dbpath string) {
	db, err := InitDatabase(dbpath)
	if err != nil {
		panic(err)
	}
	app.DB = db

	app.ui = &Ui{}

	app.ui.UIInit(500, 300)

}

func (app *App) Run() {

	// check amount of projects.
	projects, err := GetAllProjects(app.DB)
	if err != nil {
		log.Fatal(nil)
	}
	if len(projects) == 0 {
		// show timer
		app.ui.UICreateProjectWindow()
	} else {
		// show create project window
		app.ui.UICreateTimerWindow()
	}
}
