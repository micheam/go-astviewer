package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ShowError(w fyne.Window, err error) {
	dialog.NewError(err, w)
}
