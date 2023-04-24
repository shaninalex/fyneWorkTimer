package main

import (
	"database/sql"

	"fyne.io/fyne"
)

type TimerApp struct {
	tasks       []Task
	project     Project
	DB          *sql.DB
	Application fyne.App
	Window      fyne.Window
}
