package main

import (
	"net/http"
	"text/template"
)

var t = template.Must(template.ParseGlob("templates/*"))

func handler(w http.ResponseWriter, r *http.Request) {
	projects := getProjects()
	w.Header().Set("Content-Type", "text/html")
	err := t.ExecuteTemplate(w, "index", projects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}
