package main

import (
	"fmt"
	"golang_projects/photogallery_app/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeViews    *views.View
	contactViews *views.View
	faqViews     *views.View
	signUpViews  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeViews.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactViews.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqViews.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signUpViews.Render(w, nil))
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
	faqViews = views.NewView("bootstrap", "views/faq.gohtml")
	signUpViews = views.NewView("bootstrap", "views/signup.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/faq", faq)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
