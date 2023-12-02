package main

import "fmt"

func main() {
	ch := make(chan chan struct{})
	fmt.Println("Hello WebAssembly from Golang")
	<-ch
}