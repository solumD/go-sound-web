package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//парсим html-шаблон
	tmpl, err := template.ParseFiles("ui/html/home_page.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	//записываем шаблон в ResponceWriter w
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

func regGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/register.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

func regPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	w.Write([]byte(username))
	w.Write([]byte(email))
	w.Write([]byte(password))
	fmt.Println(w, r.Form)
}
