package main

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shaninalex/fyneWorkTimer/src/models"
)

func main() {

	newMode := false
	windowSize := fyne.NewSize(500, 300)

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
		time_estimate INTEGER,
	    FOREIGN KEY (project_id)
			REFERENCES projects (id)
	);
	CREATE TABLE IF NOT EXISTS timepointes (
		id INTEGER PRIMARY KEY,
		task_id INTEGER NOT NULL,
		start_time INTEGER,
		end_time INTEGER,
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

	projects, err := models.GetAllProjects(db)
	if err != nil {
		log.Fatal(err)
	}

	title := "FyneWorkTimer"

	if len(projects) == 0 {
		newMode = true
		title = "Create Project | FyneWorkTimer"
		windowSize = fyne.NewSize(300, 100)
	} else {
		newMode = false
	}

	fmt.Println(newMode)
	// TODO: check if any projects in database
	// if exists
	// show projects main window
	// else show create ProjectWindow

	Application := app.New()
	mainWindow := Application.NewWindow(title)
	mainWindow.Resize(windowSize)

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
		project := models.Project{
			Name: input.Text,
		}
		project.Create(db)
		input.Text = ""
	}))

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()

}
