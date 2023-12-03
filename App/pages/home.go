package pages

import "stackoverclash/App/components"

type Home struct {
	clashs []components.Clash
}

func (h *Home) View() string {
	clash1 := components.NewClash(0, 10, 7, []string{"Angular", "React"})
	clash2 := components.NewClash(1, 120, 54, []string{"Golang", "Rust"})

	tabClash := []components.Clash{clash1, clash2}
	h.clashs = append(h.clashs, tabClash...)
	render := ""

	for _, v := range h.clashs {
		render += v.Render()
	}
	return render
}

