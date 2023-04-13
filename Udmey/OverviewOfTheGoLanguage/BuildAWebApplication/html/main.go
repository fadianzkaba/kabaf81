package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":9991"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error Parsing template:", err)
		return
	}

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)

	fmt.Println(fmt.Sprintf("Starting Application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
