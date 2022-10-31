package dock

import (
	"github.com/rivo/tview"

	"github.com/bsthun/switchbox/program/types/ctx"
)

func Page(context *ctx.Context) tview.Primitive {
	preference := GetDockPreference()

	text := tview.NewTextView().SetText("")
	hint := tview.NewTextView().SetText("[z] move item up [x] move item down [a] apply")

	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(true).SetTitle(" Dock Arrangement ")
	WriteDockTable(list, text, &preference, 0)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(list, 0, 1, true).
		AddItem(text, 1, 0, false).
		AddItem(hint, 1, 0, false)
	return flex
}
