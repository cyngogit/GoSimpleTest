package main

import (
	"log"
	"net/http"

	"example.com/hello/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(handler.TestUrl, handler.TestHandler)
	mux.HandleFunc(handler.SecondUrl, handler.SecondHandler)
	mux.HandleFunc(handler.ThirdUrl, handler.ThirdHandler)
	mux.HandleFunc(handler.FourthUrl, handler.FourthHandler)
	mux.HandleFunc(handler.FifthUrl, handler.FifthHandler)
	mux.HandleFunc(handler.SixthUrl, handler.SixthHandler)
	mux.HandleFunc(handler.SeventhUrl, handler.SeventhHandler)
	mux.HandleFunc(handler.FormUrl, handler.FormHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("-starting web-")

	err := http.ListenAndServe(":80", mux)
	log.Println(err)

}
