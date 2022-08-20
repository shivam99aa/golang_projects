package main

import (
	"fmt"
	"golang_projects/photogallery_app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Custom 404 page</h1>")
}

func main() {
	var err error

	usersController := controllers.NewUsers()
	staticController := controllers.NewStatic()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.Handle("/", staticController.HomeView).Methods("GET")
	r.Handle("/contact", staticController.ContactView).Methods("GET")
	r.Handle("/faq", staticController.FAQView).Methods("GET")
	r.Handle("/signup", usersController.NewView).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")

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
