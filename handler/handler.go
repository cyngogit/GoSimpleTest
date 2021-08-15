package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"example.com/hello/entity"
)

var TestUrl = "/"
var SecondUrl = "/second"
var ThirdUrl = "/third"
var FourthUrl = "/fourth"
var FifthUrl = "/fifth"
var SixthUrl = "/sixth"
var SeventhUrl = "/seventh"
var FormUrl = "/form"

var TestView = "index.html"
var SecondView = "second.html"
var ThirdView = "third.html"
var FourthView = "fourth.html"
var FifthView = "fifth.html"
var SixthView = "sixth.html"
var SeventhView = "seventh.html"
var FormView = "form.html"

var ErrorView = "error.html"

// this handler for viewing index or home page
func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call test handler")

	execute(TestUrl, TestView, w, r, nil)
}

// this handler for viewing sub paging
func SecondHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call Second handler")

	execute(SecondUrl, SecondView, w, r, nil)
}

// this handler for viewing query string
func ThirdHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call third handler")

	idStr := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(idStr)

	if err != nil || idNumb < 1 {
		log.Println(err)
		execute(ThirdUrl, ErrorView, w, r, nil)
		return
	}
	idMap := map[string]string{
		"id": idStr,
	}
	execute(ThirdUrl, ThirdView, w, r, idMap)
}

// this handler for passing data to views using map
func FourthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call fourth handler")

	data := map[string]string{
		"title":   "this is a title",
		"content": "this is a content",
	}

	execute(FourthUrl, FourthView, w, r, data)
}

// this handler for passing data to views using entity
func FifthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call fifth handler")

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 22000000, Stock: 3},
		{ID: 2, Name: "Xpander", Price: 27000000, Stock: 2},
		{ID: 3, Name: "Pajero Sport", Price: 55000000, Stock: 1},
	}
	
	if (r.Method == "POST") {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "please contact our administrator", http.StatusInternalServerError)
			return
		}

		anID,_ := strconv.Atoi(r.Form.Get("ID"))
		aName := r.Form.Get("Name")
		aPrice, _ := strconv.Atoi(r.Form.Get("Price"))
		aStock, _ := strconv.Atoi(r.Form.Get("Stock"))

		anProduct := entity.Product{
			ID: anID, Name: aName, Price: aPrice, Stock: aStock}
		
		data = append(data,anProduct)
		
	}

	execute(FifthUrl, FifthView, w, r, data)
}

// this handler for passing data to views using json
func SixthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call sixth handler")

	data := []entity.Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	json.NewEncoder(w).Encode(data)
}

func SeventhHandler(w http.ResponseWriter, r *http.Request)  {
	log.Println("call seventh handler")

	method := r.Method

	switch method {
	case "GET" :
		w.Write([]byte("This method is GET"))
	case "POST" :
		w.Write([]byte("This method is POST"))
	default:
		http.Error(w, "Error is happening", http.StatusBadRequest)
		
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		execute(FormUrl,FormView,w,r,nil)
	}
}

func execute(sUrl string, sHtml string, w http.ResponseWriter, r *http.Request, data interface{}) {

	if r.URL.Path != sUrl {
		http.NotFound(w, r)
		return
	}

	sView := "views"
	sLayout := "layout.html"

	tpl, err := template.ParseFiles(path.Join(sView, sHtml), path.Join(sView, sLayout))

	if err != nil {
		log.Println(err)
		http.Error(w, "please contact our administrator", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "please contact our administrator", http.StatusInternalServerError)
		return
	}

}
