package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var app = App{}

	app.Init("timer.db")

	// var version string
	// err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(version)

	// projects, err := models.GetAllProjects(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// title := "FyneWorkTimer"

	// if len(projects) == 0 {
	// 	newMode = true
	// 	title = "Create Project | FyneWorkTimer"
	// 	windowSize = fyne.NewSize(300, 100)
	// } else {
	// 	newMode = false
	// }

	// fmt.Println(newMode)
	// // TODO: check if any projects in database
	// // if exists
	// // show projects main window
	// // else show create ProjectWindow

	// Application := app.New()
	// mainWindow := Application.NewWindow(title)
	// mainWindow.Resize(windowSize)

	// input := widget.NewEntry()
	// input.SetPlaceHolder("Enter text...")

	// content := container.NewVBox(input, widget.NewButton("Save", func() {
	// 	log.Println("Content was:", input.Text)
	// 	project := models.Project{
	// 		Name: input.Text,
	// 	}
	// 	project.Create(db)
	// 	input.Text = ""
	// }))

	// mainWindow.SetContent(content)
	// mainWindow.ShowAndRun()

}
