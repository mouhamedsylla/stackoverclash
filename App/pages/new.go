package pages

import (
	"stackoverclash/App/components"
	framego "stackoverclash/frameGo"
)

type New struct {
	//clashs []components.Clash
}

func (n *New) View() []framego.View {
	principal := components.NewPrincipale("Clash⚡", "Latest clashs")
	principaleRender := framego.NewView("#clash-title", principal.Render())
	clashRender := framego.NewView("#clash", "")
	var views []framego.View
	views = append(views, principaleRender, clashRender)
	return views
}
