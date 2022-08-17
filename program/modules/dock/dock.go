package dock

import (
	"github.com/rivo/tview"

	"github.com/bsthun/switchbox/program/types/ctx"
)

func Page(context *ctx.Context) tview.Primitive {

	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	return flex
}
