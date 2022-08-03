package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my site </h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch please send an email to <a href=\"mailto:shivam99aa@gmail.com\">shivam99aa@gmail.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404, Page not Found</h1>")
	}
	
}

func main()  {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}