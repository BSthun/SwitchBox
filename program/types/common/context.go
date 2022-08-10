package common

import "github.com/rivo/tview"

type Context struct {
	App         *tview.Application
	Pages       *tview.Pages
	CurrentPage string
}
