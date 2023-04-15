package main

import "fyne.io/fyne"

type Ui struct {
	size fyne.Size

	window fyne.Window
}

func (ui *Ui) init(width, height int) {
	ui.size = fyne.NewSize(width, height)
}

func (ui *Ui) CreateProjectWindow() {

}
