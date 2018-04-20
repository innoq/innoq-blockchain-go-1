package main

import (
	"html/template"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, "")
}
