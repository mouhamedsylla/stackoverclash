package pages

import (
	"stackoverclash/App/components"
	framego "stackoverclash/frameGo"
)

type Home struct {
	clashs []components.Clash
}

func (h *Home) View() []framego.View {
	clash1 := components.NewClash(0, 10, 7, "Angular", "React")
	clash2 := components.NewClash(1, 120, 54, "Golang", "Rust")

	tabClash := []components.Clash{clash1, clash2}
	h.clashs = append(h.clashs, tabClash...)
	render := ""

	for _, v := range h.clashs {
		render += v.Render()
	}
	principal := components.NewPrincipale("Clash⚡", "Last created clash")

	clashRender := framego.NewView("#clash", render)
	principaleRender := framego.NewView("#clash-title", principal.Render())

	var views []framego.View
	views = append(views, clashRender, principaleRender)

	return views
}

