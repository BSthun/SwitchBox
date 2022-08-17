package ctx

import "github.com/rivo/tview"

type Context struct {
	App         *tview.Application
	Pages       *tview.Pages
	CurrentPage string
}

func (r *Context) Navigate(page string) {
	r.Pages.SwitchToPage(page)
	r.CurrentPage = page
}
