package main

import (
	"log"
	"net/http"
)

// пользователь и мапа со всеми пользователями
var user = ""
var users = make(map[string]string)

// начальная страница, если пользователь не зарегистрировался
func UnSignedHandler(w http.ResponseWriter, r *http.Request) {
	//парсим html-шаблон
	if len(user) == 0 {
		//записываем шаблон в ResponseWriter w
		err := ParseNil(w, "ui/html/home_unsigned.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", 500)
		}
	} else {
		redirectToHome(w, r)
	}
}

// домашняя страница зарегистрированного пользователя
func getHomeWithSignedIn(w http.ResponseWriter, r *http.Request) {
	if len(user) == 0 {
		redirectToSignIn(w, r)
	} else {
		err := Parse(w, "ui/html/home_registered.html", user)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", 500)
		}
	}
}

// форма с регистрацией
func regGet(w http.ResponseWriter, r *http.Request) {
	err := ParseNil(w, "ui/html/register.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

func checkReg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	_, exist := users[username]
	if !exist {
		users[username] = password
		user = username //устанавливаем глобальное значение юзера
		redirectToHome(w, r)
	} else {
		redirectToLogin(w, r)
	}
}

// форма с логином
func loginGet(w http.ResponseWriter, r *http.Request) {
	err := ParseNil(w, "ui/html/login.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	us := r.PostFormValue("username")
	pass := r.PostFormValue("password")
	if password, exist := users[us]; exist {
		if pass == password {
			user = us //устанавливаем глобальное значение юзера
			redirectToHome(w, r)
		} else {
			http.Error(w, "Invalable password", http.StatusBadRequest)
		}
	} else {
		redirectToSignIn(w, r)
	}
}

// выход пользователя
func Exit(w http.ResponseWriter, r *http.Request) {
	user = ""
	redirectToUnregistered(w, r)
}
