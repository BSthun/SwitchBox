package dashboard

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StatusLinePrimitive() tview.Primitive {
	return tview.NewBox().SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		flex := tview.NewFlex().SetDirection(tview.FlexColumn)

		if width > 20 {
			flex.AddItem(tview.NewBox().SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
				tview.Print(screen, "bsthun@thun-macmini", x, y, width, tview.AlignCenter, tcell.ColorLightCyan)
				return 0, 0, 0, 0
			}), 0, 1, false)
		}
		if width > 40 {
			flex.AddItem(tview.NewBox().SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
				tview.Print(screen, "10.2.1.107", x, y, width, tview.AlignCenter, tcell.ColorLightCyan)
				return 0, 0, 0, 0
			}), 0, 1, false)
		}
		if width > 60 {
			flex.AddItem(tview.NewBox().SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
				tview.Print(screen, "12% CPU (2.6GHz)", x, y, width, tview.AlignCenter, tcell.ColorLightCyan)
				return 0, 0, 0, 0
			}), 0, 1, false)
		}
		if width > 80 {
			flex.AddItem(tview.NewBox().SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
				tview.Print(screen, "24498M used / 32G", x, y, width, tview.AlignCenter, tcell.ColorLightCyan)
				return 0, 0, 0, 0
			}), 0, 1, false)
		}
		if width > 100 {
			flex.AddItem(tview.NewBox().SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
				tview.Print(screen, "Load 3.64 2.63 2.49", x, y, width, tview.AlignCenter, tcell.ColorLightCyan)
				return 0, 0, 0, 0
			}), 0, 1, false)
		}

		flex.SetRect(x, y, width, height)
		flex.Draw(screen)

		return 0, 0, 0, 0
	})
}
