package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
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

func (ui *Ui) UICreateProjectWindow() {

}

func (ui *Ui) UICreateTimerWindow() {

}
