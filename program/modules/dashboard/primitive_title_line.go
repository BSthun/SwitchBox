package dashboard

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TitleLinePrimitive() tview.Primitive {
	prim := tview.NewBox().SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		// * Draw a horizontal line across the middle of the box.
		for cx := x + 1; cx < x+width-1; cx++ {
			screen.SetContent(cx, 0, tview.BoxDrawingsLightHorizontal, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
		}

		// * Write title at center of line.
		tview.Print(screen, " SwitchBox ", x+1, y, width, tview.AlignCenter, tcell.ColorYellow)

		// Space for other content.
		return 0, 0, 0, 0
	})
	return prim
}
