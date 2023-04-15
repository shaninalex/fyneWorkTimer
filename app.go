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
		// panic(err)
		log.Println(err)
	}
	app.DB = db
	log.Println("db initialized")

	app.ui = &Ui{}
	app.ui.UIInit(500, 300)
	log.Println("ui initialized")
}

func (app *App) Run() {

	log.Println("Application started")
	// check amount of projects.
	projects, err := GetAllProjects(app.DB)
	if err != nil {
		log.Fatal(nil)
	}
	if len(projects) == 0 {
		// show timer
		app.ui.UICreateProjectWindow(app.DB)
	} else {
		// show create project window
		app.ui.UICreateTimerWindow()
	}

	app.ui.Run()
	defer app.DB.Close()
}
