package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type Ui struct {
	Size        fyne.Size
	Application fyne.App
	Window      fyne.Window
}

func (ui *Ui) UIInit(width, height int) {
	ui.Size = fyne.NewSize(width, height)
	ui.Application = app.New()
	ui.Window = ui.Application.NewWindow("hello")
	ui.Window.Resize(ui.Size)
}

func (ui *Ui) Run() {
	ui.Window.ShowAndRun()
}

func (ui *Ui) UICreateProjectWindow(db *sql.DB) {
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := container.NewVBox(input, widget.NewButton("Save", func() {
		// TODO: move this func out of this into separate handlers
		log.Println("Project Name:", input.Text)
		// TODO: save to database and switch to Timer Window
		project := &Project{}
		project.Name = input.Text
		project.Create(db)
		input.Text = ""
		ui.UICreateTimerWindow()
	}))
	ui.Window.SetContent(content)
}

func (ui *Ui) UICreateTimerWindow() {
	label := widget.NewLabel("Main window")
	content := container.NewVBox(label)
	ui.resize(600, 500)
	ui.Window.SetContent(content)
}

func (ui *Ui) resize(width, height int) {
	ui.Size = fyne.NewSize(width, height)
	ui.Window.Resize(ui.Size)
}
