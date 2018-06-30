package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", indexHandle)

	http.HandleFunc("/get", getHandle)
	http.HandleFunc("/post", postHandle)
	http.HandleFunc("/add", addHandle)
}

// var templates = template.Must(template.ParseFiles("index.html"))
// var validPath = regexp.MustCompile("^/(index)/([a-zA-Z0-9]+)$")

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	err := templates.ExecuteTemplate(w, tmpl+".html", "halo")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func indexHandle(w http.ResponseWriter, r *http.Request) {

	// renderTemplate(w, "index")

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func getHandle(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from todo;")
	if err != nil {
		log.Panic(err)
	}

	list := get_all(rows)

	t, _ := template.ParseFiles("templates/get.html")
	t.Execute(w, list)
}

func addHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/post.html")
	t.Execute(w, nil)
}

func postHandle(w http.ResponseWriter, r *http.Request) {
	// stmt, err := db.Prepare("INSERT todo SET title=?,finished=?,created=?")
	// if err != nil {
	// 	log.Panic(err)
	// }

	// res, err := stmt.Exec("reading", false, time.Now())
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(res.LastInsertId())

	title := r.FormValue("title")
	finished := r.FormValue("finished")
	fmt.Println(title, finished)

	http.Redirect(w, r, "/get", http.StatusFound)
}

type SignInResponse struct {
	UserID       string `json:"user_id"`
	SessionToken string `json:"session_token"`
}
