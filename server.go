package main

import (
	"net/http"
	"text/template"
)

var tmplts = template.Must(template.ParseGlob("templates/*"))

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	projects := getProjects()
	err := tmplts.ExecuteTemplate(w, "index", projects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := tmplts.ExecuteTemplate(w, "about", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// For Openshift Deployment
	// bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))

	// For local testing
	bind := ":8080"

	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/about/", aboutPageHandler)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(bind, nil)
}
