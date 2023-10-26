package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	appID          = "github.com/micheam/go-astviewer"
	title          = "Go AST Viewer"
	srcPlaceholder = `package main

import (
    "fmt"
    "os"
)

func main() {
    msg := "hello, world"
    fmt.Fprintln(os.Stdout, msg)
}`
)

var (
	windowSize = fyne.NewSize(900., 400.)

	inputData = binding.NewString()
	srcPane   fyne.CanvasObject

	outputData = binding.NewString()
	destPane   fyne.CanvasObject
)

func main() {
	a := app.NewWithID(appID)
	// TODO: a.SetIcon
	a.Settings().SetTheme(&mytheme{})
	w := a.NewWindow("Go AST Viewer")
	w.Resize(windowSize)
	w.SetContent(loadui(w))
	w.ShowAndRun()
}

func loadui(w fyne.Window) fyne.CanvasObject {
	var top, bottom, left, right, content fyne.CanvasObject
	initContent(w)
	// TODO: toggle button to switch layout
	content = container.NewStack(container.NewHSplit(srcPane, destPane))
	return container.NewBorder(top, bottom, left, right, content)
}

func initContent(w fyne.Window) {
	inputData.Set(srcPlaceholder)
	entry := widget.NewEntryWithData(inputData)
	entry.MultiLine = true
	entry.Wrapping = fyne.TextWrapWord
	entry.TextStyle = fyne.TextStyle{Monospace: true}
	entry.OnChanged = func(s string) {
		astTree, err := parse(s)
		if err != nil {
			ShowError(w, fmt.Errorf("parse: %s", err))
			return
		}
		_ = outputData.Set(astTree)
	}
	srcPane = container.New(layout.NewPaddedLayout(), entry)

	result := widget.NewEntryWithData(outputData)
	result.MultiLine = true
	result.Wrapping = fyne.TextWrapWord
	result.TextStyle = fyne.TextStyle{Monospace: true}
	result.Disable()
	destPane = container.New(layout.NewPaddedLayout(), result)

	{ // handle initial data
		d, _ := inputData.Get()
		astTree, _ := parse(d)
		_ = outputData.Set(astTree)
	}
}

func parse(src string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, parser.Mode(0))
	if err != nil {
		return "", err
	}
	wr := new(strings.Builder)
	ast.Fprint(wr, fset, f, ast.NotNilFilter)
	return wr.String(), nil
}
