package main

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "timer.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// define schema
	scheme := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY,
		name VARCHAR(64) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		name VARCHAR(64),
		project_id INTEGER NOT NULL,
	    FOREIGN KEY (project_id)
			REFERENCES projects (id)
	);
	CREATE TABLE IF NOT EXISTS timepointes (
		id INTEGER PRIMARY KEY,
		task_id INTEGER NOT NULL,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL,
	    FOREIGN KEY (task_id)
			REFERENCES task (id)	
	)
	`
	_, err = db.Exec(scheme)
	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)

	myApp := app.New()
	myWindow := myApp.NewWindow("FyneWorkTimer")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
		input.Text = ""
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
