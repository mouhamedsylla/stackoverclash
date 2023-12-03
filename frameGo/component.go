package framego

import (
	"bytes"
	"fmt"
	"html/template"
)

// The above code defines an interface named "Component" in the Go programming language.
// @property {string} Render - The `Render` method is a function that returns a string. It is used to
// generate the HTML or UI representation of a component.
type Component interface {
	Render() string
}

// The function `ParseComponent` takes a string and an interface as input, parses the string as a
// template, executes the template with the provided interface, and returns the result as a string.
func ParseComponent(c string, p interface{}) string{
	var resultBuffer bytes.Buffer
	tpml, err1 := template.New("template").Parse(c)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	err := tpml.Execute(&resultBuffer, &p)
	if err != nil {
		panic(err)
	}
	return resultBuffer.String()
}
