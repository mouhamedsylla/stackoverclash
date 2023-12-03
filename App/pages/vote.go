package pages

import (
	"stackoverclash/App/components"
	framego "stackoverclash/frameGo"
)

type Vote struct {
	//clashs []components.Clash
}

func (n *Vote) View() []framego.View {
	principal := components.NewPrincipale("Clash⚡", "Vote to clash !")
	principaleRender := framego.NewView("#clash-title", principal.Render())
	clashRender := framego.NewView("#clash", "")
	var views []framego.View
	views = append(views, principaleRender, clashRender)
	return views
}