package tview

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run() {
	app := tview.NewApplication()

	boxTop := tview.NewBox().SetBorder(true).SetTitle("Top").SetBackgroundColor(tcell.ColorDefault)
	boxBottom := tview.NewBox().SetBorder(true).SetTitle("Bottom").SetBackgroundColor(tcell.ColorDefault)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(boxTop, 10, 1, false).
		AddItem(boxBottom, 0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
