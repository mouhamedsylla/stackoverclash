package main

import (
	"net/http"
	"log"
)

const port = ":8081"

func main() {
	static := http.FileServer(http.Dir("../assets/"))
	log.Printf("Server started on http://localhost%s ...", port)
	http.ListenAndServe(port, static)
}