package components

import (
	framego "stackoverclash/frameGo"
)

type Principale struct {
	Mode  string
	Title string
	With string
}

func (p *Principale) Render() string {
	component :=
		`
	<div style="width: {{.With}}em;" id="clash-label">{{.Mode}}</div>
    <h2 id="titre2">{{.Title}}</h2>
	`
	return framego.ParseComponent(component, &p)
}

func NewPrincipale(m, t, w string) Principale {
	return Principale{
		Mode: m,
		Title: t,
		With: w,
	}
}
