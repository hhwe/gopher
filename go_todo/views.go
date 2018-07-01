package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func init() {
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delte", deleteHandler)

	http.Handle("/templates/", http.StripPrefix("/templates/",
		http.FileServer(http.Dir("./templates"))))
}

// var templates = template.Must(template.ParseFiles("index.html"))
// var validPath = regexp.MustCompile("^/(index)/([a-zA-Z0-9]+)$")

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	err := templates.ExecuteTemplate(w, tmpl+".html", "halo")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// renderTemplate(w, "index")

	t, _ := template.ParseFiles("templates/layout.html", "templates/index.html")
	t.Execute(w, nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from todo;")
	if err != nil {
		log.Panic(err)
	}

	list := getAll(rows)

	t, _ := template.ParseFiles("templates/layout.html", "templates/get.html")
	t.Execute(w, list)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/layout.html", "templates/post.html")
	t.Execute(w, nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "title is needed", http.StatusInternalServerError)
		return
	}
	finished, err := strconv.ParseBool(r.FormValue("finished"))
	if err != nil {
		log.Panic(err)
	}
	stmt, err := db.Prepare("INSERT todo SET title=?,finished=?,created=?")
	if err != nil {
		log.Panic(err)
	}

	res, err := stmt.Exec(title, finished, time.Now())
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res.LastInsertId())

	fmt.Println(title, finished)

	http.Redirect(w, r, "/get", http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 10)
	if err != nil {
		log.Panic(err)
	}

	stmt, err := db.Prepare("delete from todo where id=?")
	if err != nil {
		log.Panic(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res.LastInsertId())

	http.Redirect(w, r, "/get", http.StatusFound)
}

// type SignInResponse struct {
// 	UserID       string `json:"user_id"`
// 	SessionToken string `json:"session_token"`
// }
