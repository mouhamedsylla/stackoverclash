package framego

import (
	"bytes"
	"fmt"
	"html/template"
)

type Component interface {
	Render() string
}

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
