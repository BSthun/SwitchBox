package dashboard

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/bsthun/switchbox/program/types/common"
)

//go:embed embed_music_osascript.scpt
var NowPlayingMusicOsaScript string

func NowPlayingPrimitive(context *common.Context) tview.Primitive {
	musics := NowPlayingGetMusic()
	return tview.NewBox().SetBorder(true).SetTitle(" Now Playing ").SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		flex := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(NowPlayingWavePrimitive(context), 5, 0, false).
			AddItem(tview.NewTextView().SetTextColor(tcell.ColorLightGreen).SetText(" ♪ "+musics[0]), 1, 0, false).
			AddItem(tview.NewTextView().SetTextColor(tcell.ColorLightSalmon).SetText(" ★ "+musics[1]), 1, 0, false).
			AddItem(tview.NewTextView().SetTextColor(tcell.ColorLightCoral).SetText(" ❏ "+musics[2]), 1, 0, false).
			AddItem(tview.NewTextView().SetTextColor(tcell.ColorLightSeaGreen).SetText(" ⦿ "+musics[3]), 1, 0, false).
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
			for last+change <= 0 || last+change >= height {
				change = rand.Intn(3) - 1
			}
			waves = waves[1:]
			waves = append(waves, last+change)

			for i, wave := range waves {
				for j := 0; j < wave; j++ {
					ch := '▇'
					if j == wave-1 {
						r := rand.Intn(4)
						if r == 1 {
							ch = '▃'
						}
						if r == 2 {
							ch = '▅'
						}
						if r == 3 {
							ch = '▆'
						}
					}
					screen.SetContent(x+i, height+y-j-2, ch, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
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
			NowPlayingWaveTicker = time.NewTicker(time.Millisecond * 300)
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

func NowPlayingGetMusic() [4]string {
	for _, player := range []string{"Spotify", "Music"} {
		osa := strings.ReplaceAll(NowPlayingMusicOsaScript, "{{PLAYER}}", player)
		result, err := exec.Command("osascript", "-e", osa).Output()

		if err != nil || len(result) == 0 {
			continue
		}

		split := strings.Split(string(result), " ः ")

		duration, _ := strconv.ParseInt(split[3], 10, 64)
		position, _ := strconv.ParseFloat(split[4], 32)
		return [4]string{split[0], split[1], split[2], fmt.Sprintf("%.0f:%02d / %d:%02d", position/60, int(position)%60, duration/1000/60, duration/1000%60)}
	}

	return [4]string{}
}
