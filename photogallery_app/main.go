package main

import (
	"fmt"
	"golang_projects/photogallery_app/controllers"
	"golang_projects/photogallery_app/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeViews    *views.View
	contactViews *views.View
	faqViews     *views.View
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
	usersController := controllers.NewUsers()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	r.HandleFunc("/faq", faq).Methods("GET")

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
