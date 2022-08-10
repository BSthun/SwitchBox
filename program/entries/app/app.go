package app

import (
	"github.com/rivo/tview"

	"github.com/bsthun/switchbox/program/modules/dashboard"
	"github.com/bsthun/switchbox/program/types/common"
	"github.com/bsthun/switchbox/program/wrappers/interactive"
)

var App *tview.Application

func Entrypoint() {
	// * Initialize app
	App = tview.NewApplication()

	// * Initialize page
	pages := tview.NewPages()

	// * Initialize context
	context := &common.Context{
		App:         App,
		Pages:       pages,
		CurrentPage: "list",
	}

	pages.AddPage("dashboard", dashboard.Page(context), true, true)

	list := tview.NewList().
		AddItem("Dashboard", "System dashboard", '1', func() {
			context.CurrentPage = "dashboard"
			pages.SwitchToPage("dashboard")
		}).
		AddItem("Dock", "Manage macOS Dock arrangement", '2', nil).
		AddItem("Quit", "", 'q', func() {
			App.Stop()
		})
	pages.AddPage("list", list, true, true)

	// * Launch app
	App.SetRoot(pages, true)

	if err := App.Run(); err != nil {
		interactive.Throw("app", "error running app", err, nil)
	}
}
