package main

import (
	"fmt"
	"stackoverclash/App/pages"
	framego "stackoverclash/frameGo"
	"syscall/js"
)

func main() {
	ch := make(chan chan struct{})
	fmt.Println("Hello WebAssembly from Golang")

	router := &framego.Router{}
	routes := []string{
		"/",
		"/new",
		"/popular",
		"/vote",
		"/create",
	}
	pages := []pages.Pages{
		&pages.Home{},
		&pages.New{},
		&pages.Popular{},
		&pages.Vote{},
		&pages.Createclash{},
	}
	for i, r := range routes {
		router.AddRoute(r, pages[i].View())
	}
	home := js.Global().Get("location").Get("href").String()
	router.ResolvRoute(home)
	elements := js.Global().Get("document").Call("querySelectorAll", "a")
	for i := 0; i < elements.Length(); i++ {
		element := elements.Index(i)
		element.Call("addEventListener", "click", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
			event := p[0]
			event.Call("preventDefault")
			url := element.Get("href").String()
			router.ResolvRoute(url)
			return nil
		}))
	}
	<-ch
}
