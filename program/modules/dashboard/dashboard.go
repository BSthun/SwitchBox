package dashboard

import (
	"github.com/rivo/tview"

	"github.com/bsthun/switchbox/program/types/common"
)

func Page(context *common.Context) tview.Primitive {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(TitleLinePrimitive(), 2, 0, false).
		AddItem(StatusLinePrimitive(), 2, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(NowPlayingPrimitive(context), 30, 0, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 0, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 0, 1, false)
	return flex
}
