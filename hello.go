package main

import (
	"log"
	"net/http"

	"example.com/hello/route"
)

func main() {
	mux := http.NewServeMux()

	route.SetupLink(mux)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("-starting web-")

	err := http.ListenAndServe(":80", mux)
	log.Println(err)

}
