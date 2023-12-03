package pages


import (
	"stackoverclash/App/components"
	framego "stackoverclash/frameGo"
)

type Createclash struct {
	//clashs []components.Clash
}

func (c *Createclash) View() []framego.View {
	principal := components.NewPrincipale("Create clash", "New clash ✨", "8")
	principaleRender := framego.NewView("#clash-title", principal.Render())
	clashRender := framego.NewView("#clash", "")
	var views []framego.View
	views = append(views, principaleRender, clashRender)
	return views
}