package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := chi.NewRouter()

	//главная страница без выполненного входа
	r.Get(`/`, UnSignedHandler)

	//главная страница с выполненным входом
	r.Get(`/home`, getHomeWithSignedIn)

	//регистрация пользователя и проверка есть ли такой пользователь
	r.Get(`/signin`, regGet)
	r.Post(`/checksignin`, checkReg)

	//вход в существующий аккаунт и проверка есть ли такой
	r.Get(`/login`, loginGet)
	r.Post(`/checklogin`, checkLogin)

	//выход пользователя
	r.Get(`/unsign`, Exit)

	//получение своих избранных книг
	r.Get("/favorite", favoriteGet)

	//страница с избранными книгами для не зарегистрированного пользователя
	r.Get("/favorite/unsigned", favoriteUnSignedGet)

	//преобразуем визуал
	fs := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	db, err := OpenDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
