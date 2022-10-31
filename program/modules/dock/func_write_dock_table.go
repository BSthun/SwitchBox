package dock

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/bsthun/switchbox/program/types/extern"
)

func WriteDockTable(list *tview.List, text *tview.TextView, preference *extern.DockPreference, state int) {
	for i, item := range preference.RecentApps {
		list.AddItem(item.TileData.FileLabel, "", rune(strconv.Itoa(i + 1)[0]), nil)
	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		direction := "none"

		if event.Name() == "Rune[a]" {
			ApplyDock(preference)
		}

		if event.Name() == "Rune[z]" || event.Name() == "Rune[x]" {
			if event.Name() == "Rune[z]" {
				if state == 0 {
					return nil
				}
				preference.RecentApps[state], preference.RecentApps[state-1] = preference.RecentApps[state-1], preference.RecentApps[state]
				direction = "up"
			}
			if event.Name() == "Rune[x]" {
				if state == len(preference.RecentApps)-1 {
					return nil
				}
				preference.RecentApps[state], preference.RecentApps[state+1] = preference.RecentApps[state+1], preference.RecentApps[state]
				direction = "down"
			}
		}

		if (event.Name() == "Up" && state > 0) || direction == "up" {
			state--
		}
		if (event.Name() == "Down" && state < len(preference.RecentApps)-1) || direction == "down" {
			state++
		}

		if direction != "none" {
			list.Clear()
			WriteDockTable(list, text, preference, state)
		}

		text.SetText(fmt.Sprintf("Selecting #%d %s", state, preference.RecentApps[state].TileData.BundleIdentifier))
		list.SetCurrentItem(state)

		return nil
	})

	text.SetText(fmt.Sprintf("Selecting #%d %s", state+1, preference.RecentApps[state].TileData.BundleIdentifier))
}
