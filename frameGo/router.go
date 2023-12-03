package framego

import (
	"regexp"
	"syscall/js"
)

// The Route type represents a route in a Go web application, consisting of a path and a list of views.
// @property {string} path - The `path` property in the `Route` struct represents the URL path for the
// route. It specifies the URL pattern that the route will match against.
// @property {[]View} views - The `views` property is a slice of `View` objects. It represents the
// different views or templates that can be rendered for this route.
type Route struct {
	path  string
	views []View
}

// The `Router` type represents a router in Go.
// @property {[]Route} routes - The `routes` property is a slice of `Route` structs. It is used to
// store the different routes that the router will handle.
type Router struct {
	routes []Route
}

// The `AddRoute` function is a method of the `Router` struct. It is used to add a new route to the
// router's list of routes.
func (r *Router) AddRoute(p string, v []View) {
	myRoute := Route{
		path:  p,
		views: v,
	}
	r.routes = append(r.routes, myRoute)
}


// The `ResolvRoute` function is a method of the `Router` struct. It is used to resolve a route based
// on the given URL and update the HTML content of the corresponding view.
func (r *Router) ResolvRoute(URL string) {
	var element js.Value
	match := MatchPath(URL)
	for i, route := range r.routes {
		if route.path == match {
			for _, view := range route.views {
				element = js.Global().Get("document").Call("querySelector", view.selector)
				element.Set("innerHTML", view.html)
				js.Global().Get("history").Call("pushState", js.Null(), "", route.path)
			}
			break
		} else if i == len(r.routes)-1 {
			element.Set("innerHTML", "<h1>Error 404: Page Not Found")
		}
	}
}

// The MatchPath function extracts the path from a given URL.
func MatchPath(url string) string {
	pathRegex := regexp.MustCompile(`http://[^/]+(/.*)?`)
	match := pathRegex.FindStringSubmatch(url)
	if match != nil {
		return match[1]
	}
	return ""
}


