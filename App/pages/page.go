package pages

import framego "stackoverclash/frameGo"

type Pages interface{
	View() []framego.View
}