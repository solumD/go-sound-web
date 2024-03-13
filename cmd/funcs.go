package main

import (
	"html/template"
	"log"
	"net/http"
)

// парсинг шаблона с параметрами
func Parse(w http.ResponseWriter, path string, params string) error {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return err
	}
	//записываем шаблон в ResponseWriter w
	err = tmpl.Execute(w, params)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return err
	}
	return nil
}

// парсинг шаблона без параметров
func ParseNil(w http.ResponseWriter, path string) error {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return err
	}
	//записываем шаблон в ResponseWriter w
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return err
	}
	return nil
}
