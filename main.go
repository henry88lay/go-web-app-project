package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w, http.ResponseWriter, r *http.Request) {
	w.Header().set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Let's get in touch, please send an email to <a href=\"mailto:support@photobros.com\">support@photobros.com</a>.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
