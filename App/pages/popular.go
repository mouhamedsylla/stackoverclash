package pages

import (
	"stackoverclash/App/components"
	framego "stackoverclash/frameGo"
)

type Popular struct {
	//clashs []components.Clash
}

func (p *Popular) View() []framego.View {
	principal := components.NewPrincipale("Clash⚡", "Popular clash")
	principaleRender := framego.NewView("#clash-title", principal.Render())
	clashRender := framego.NewView("#clash", "")
	var views []framego.View
	views = append(views, principaleRender, clashRender)
	return views
}