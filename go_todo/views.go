package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func init() {
	http.HandleFunc("/", indexHandle)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := template.ExecuteTemplate(w, tmpl+".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, "hello")
}

type SignInResponse struct {
	UserID       string `json:"user_id"`
	SessionToken string `json:"session_token"`
}
