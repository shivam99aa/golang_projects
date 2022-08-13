package main

import (
	"fmt"
	"golang_projects/photogallery_app/views"
	"net/http"

	"github.com/gorilla/mux"
)

var homeViews *views.View
var contactViews *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeViews.Template.ExecuteTemplate(w, homeViews.Layouts, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactViews.Template.ExecuteTemplate(w, contactViews.Layouts, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "FAQ Page")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Custom 404 page</h1>")
}

func main() {
	var err error
	homeViews = views.NewView("bootstrap", "views/home.gohtml")
	contactViews = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
