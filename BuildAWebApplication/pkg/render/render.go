package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders template using the html
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error Parsing template:", err)
		return
	}

}
