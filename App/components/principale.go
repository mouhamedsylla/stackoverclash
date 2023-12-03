package components

import (
	framego "stackoverclash/frameGo"
)

type Principale struct {
	Mode  string
	Title string
}

func (p *Principale) Render() string {
	component :=
		`
	<div id="clash-label">{{.Mode}}</div>
    <h2 id="titre2">{{.Title}}</h2>
	`
	return framego.ParseComponent(component, &p)
}
