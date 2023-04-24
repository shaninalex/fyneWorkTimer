package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type WorkTimer struct {
	DB              *sql.DB
	Application     fyne.App
	Window          fyne.Window
	projects        []Project
	selectedProject Project
}

func (wt *WorkTimer) Init(dbpath string) {
	db, err := InitDatabase(dbpath)
	if err != nil {
		// panic(err)
		log.Println(err)
	}
	wt.DB = db
	log.Println("db initialized")

	size := fyne.NewSize(600, 400)
	wt.Application = app.New()
	wt.Window = wt.Application.NewWindow("hello")
	wt.Window.Resize(size)
	log.Println("ui initialized")
}

func (wt *WorkTimer) Run() {

	log.Println("Application started")
	// check amount of projects.
	projects, err := GetAllProjects(wt.DB)
	if err != nil {
		log.Fatal(nil)
	}
	if len(projects) == 0 {
		// show timer
		wt.UICreateProjectWindow()
	} else {
		// show create project window
		wt.UICreateTimerWindow()
	}

	wt.Window.ShowAndRun()
	defer wt.DB.Close()
}

func (wt *WorkTimer) UICreateProjectWindow() {
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := container.NewVBox(layout.NewSpacer(), input, widget.NewButton("Save", func() {
		// TODO: move this func out of this into separate handlers
		log.Println("Project Name:", input.Text)
		// TODO: save to database and switch to Timer Window
		project := &Project{}
		project.Name = input.Text
		project.Create(wt.DB)
		input.Text = ""
		wt.UICreateTimerWindow()
	}), layout.NewSpacer())
	wt.Window.SetContent(content)
}

func (wt *WorkTimer) UICreateTimerWindow() {
	sidebar := container.NewVBox(wt.sidebarContent())
	timerContent := container.NewVBox(wt.mainContent())

	split := container.NewHSplit(sidebar, timerContent)
	split.Offset = 0
	wt.Window.SetContent(split)
}

func (wt *WorkTimer) sidebarContent() (*widget.Select, *widget.Button) {
	projects, err := GetAllProjects(wt.DB)

	wt.projects = projects
	wt.selectedProject = projects[0]
	if err != nil {
		log.Println(err)
	}

	log.Println(projects)

	options := []string{}

	for _, element := range projects {
		options = append(options, element.Name)
	}

	selectWidget := widget.NewSelect(options, nil)
	selectWidget.Selected = options[0]
	// selectedOption := widget.NewLabel("No option selected")

	// Set the onSelect function for the select widget
	selectWidget.OnChanged = func(selected string) {
		// selectedOption.SetText("Selected option: " + selected)
	}

	buttonCreateProject := widget.NewButton("Create new project", func() {
		wt.UICreateProjectWindow()
	})

	return selectWidget, buttonCreateProject
}

func (wt *WorkTimer) mainContent() *fyne.Container {
	input := widget.NewEntry()
	input.SetPlaceHolder("New Task Name")

	projects, err := GetAllTasks(wt.DB, wt.selectedProject.Id)
	if err != nil {
		log.Println(err)
	}
	task := Task{}
	createTaskBtn := widget.NewButton("create task", func() {
		if len(input.Text) > 0 {
			task.Id = wt.selectedProject.Id
			task.Name = input.Text
			task.TimeEstimate = 1231
			task.Create(wt.DB)
			// TODO: add new item in tasks list
			projects = append(projects, task)
		} else {
			log.Println("Unable to create task without name...")
			// TODO: show error label and on start typing - hide error
		}
	})

	list := widget.NewList(
		func() int {
			return len(projects)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(projects[i].Name)
		},
	)

	createTaskRow := widget.NewHBox(input, createTaskBtn)
	mainWidgetContainer := container.NewVBox(createTaskRow, list)
	return mainWidgetContainer
}
