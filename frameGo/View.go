package framego

// The View type represents a view in a Go program, with a selector and HTML content.
// @property {string} selector - The `selector` property in the `View` struct represents a CSS selector
// that is used to identify and select elements in the HTML document. It is typically used to target
// specific elements for manipulation or styling using JavaScript or CSS.
// @property {string} html - The `html` property in the `View` struct represents the HTML content that
// will be displayed in the view. It is a string that contains the HTML markup for the view's content.
type View struct {
	selector string
	html     string
}

// The NewView function creates a new View object with the given selector and HTML content.
func NewView(s, h string) View {
	return View{
		selector: s,
		html:     h,
	}
}