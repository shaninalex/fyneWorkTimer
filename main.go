package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	// First U want to declare app layout
	// and then wirking with data.
	// var app = &App{}
	// app.Init("timer.db")
	// app.Run()
	timerApp := TimerApp{}
	timerApp.project = Project{Name: "Create Fyne App"}
	timerApp.tasks = []Task{
		{
			Id:           1,
			Name:         "Setup dev environment",
			ProjectId:    1,
			TimeEstimate: 940234,
		},
		{
			Id:           2,
			Name:         "Setup Layout",
			ProjectId:    1,
			TimeEstimate: 940234,
		},
		{
			Id:           3,
			Name:         "Sync with database",
			ProjectId:    1,
			TimeEstimate: 940234,
		},
		{
			Id:           4,
			Name:         "Create timer progress ( or actualy ) regress bar for timer",
			ProjectId:    1,
			TimeEstimate: 940234,
		},
	}
	timerApp.Application = app.New()

	w := timerApp.Application.NewWindow("Fyne Demo")

	w.SetMaster()

	content := container.NewMax()
	intro.Wrapping = fyne.TextWrapWord

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

	split := container.NewHSplit(get_task_list(), tutorial)

	split.Offset = 0
	w.SetContent(split)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func get_task_list() fyne.CanvasObject {

	tree := widget.NewTreeWithStrings(menuItems)
	// tree.OnSelected =

	return container.NewBorder(nil, nil, nil, nil, tree)
}

var menuItems = map[string][]string{
	"":            {"welcome", "collections", "advanced"},
	"collections": {"list", "table"},
}
