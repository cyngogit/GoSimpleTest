package route

import (
	"net/http"

	"example.com/hello/handler"
)

func SetupLink(mux *http.ServeMux) {
	mux.HandleFunc(handler.TestUrl, handler.TestHandler)
	mux.HandleFunc(handler.SecondUrl, handler.SecondHandler)
	mux.HandleFunc(handler.ThirdUrl, handler.ThirdHandler)
	mux.HandleFunc(handler.FourthUrl, handler.FourthHandler)
	mux.HandleFunc(handler.FifthUrl, handler.FifthHandler)
	mux.HandleFunc(handler.SixthUrl, handler.SixthHandler)
	mux.HandleFunc(handler.SeventhUrl, handler.SeventhHandler)
	mux.HandleFunc(handler.FormUrl, handler.FormHandler)

}
