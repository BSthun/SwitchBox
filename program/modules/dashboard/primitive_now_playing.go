package dashboard

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/bsthun/switchbox/program/types/common"
)

func NowPlayingPrimitive(context *common.Context) tview.Primitive {
	return tview.NewBox().SetBorder(true).SetTitle(" Now Playing ").SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		flex := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(NowPlayingWavePrimitive(context), 6, 0, false).
			AddItem(tview.NewTextView().SetTextColor(tcell.ColorLightGreen).SetText("You Don't Know"), 1, 0, false).
			AddItem(tview.NewTextView().SetTextColor(tcell.ColorLightSalmon).SetText("Katelyn Tarver"), 1, 0, false).
			AddItem(tview.NewTextView().SetTextColor(tcell.ColorLightCoral).SetText("Tired Eyes"), 1, 0, false).
			AddItem(tview.NewBox(), 0, 1, false)

		flex.SetRect(x+1, y+1, width-2, height-2)
		flex.Draw(screen)

		return 0, 0, 0, 0
	})
}

var NowPlayingWaveTicker *time.Ticker

func NowPlayingWavePrimitive(context *common.Context) tview.Primitive {
	draw := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		var waves []int
		for i := 0; i < width; i++ {
			waves = append(waves, rand.Intn(height))
		}

		f := func() {
			last := waves[len(waves)-1]
			change := rand.Intn(3) - 1
			for last+change < 0 || last+change >= height {
				change = rand.Intn(3) - 1
			}
			waves = waves[1:]
			waves = append(waves, last+change)

			for i, wave := range waves {
				for j := 0; j < wave; j++ {
					screen.SetContent(x+i, (height+y)-j-2, '█', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
				}
				for j := wave; j < height-1; j++ {
					screen.SetContent(x+i, (height+y)-j-2, ' ', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
				}
			}
		}

		if context.CurrentPage != "dashboard" {
			return 0, 0, 0, 0
		}

		f()
		if NowPlayingWaveTicker == nil {
			NowPlayingWaveTicker = time.NewTicker(time.Millisecond * 500)
			go func() {
				for range NowPlayingWaveTicker.C {
					f()
					screen.Sync()
				}
			}()
		}

		return 0, 0, 0, 0
	}

	box := tview.NewBox().SetDrawFunc(draw)

	return box
}
