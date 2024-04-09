package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// пользователь и мапа со всеми пользователями
var currentUser = ""
var currentUserId = 0

// начальная страница, если пользователь не зарегистрировался
func UnSignedHandler(w http.ResponseWriter, r *http.Request) {
	//парсим html-шаблон
	if len(currentUser) == 0 {
		//записываем шаблон в ResponseWriter w
		err := ParseNil(w, "ui/html/home_unsigned.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", 500)
		}
	} else {
		redirectToHome(w, r)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

// домашняя страница зарегистрированного пользователя
func getHomeWithSignedIn(w http.ResponseWriter, r *http.Request) {
	if len(currentUser) == 0 {
		redirectToSignIn(w, r)
	} else {
		err := Parse(w, "ui/html/home_registered.html", currentUser)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", 500)
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

// форма с регистрацией
func regGet(w http.ResponseWriter, r *http.Request) {
	err := ParseNil(w, "ui/html/register.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func checkReg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//открываем базу данных
	db, err := OpenDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	//отправляем select-запрос
	_, _, err = selectUser(db, login)
	if err == sql.ErrNoRows {
		//добавляем пользователя
		err := insertUser(db, login, password)
		if err != nil {
			log.Println(err)
		}
		id, _, err := selectUser(db, login)
		if err != nil {
			log.Println(err)
		}
		currentUser = login //устанавливаем глобальное значение юзера
		currentUserId = id
		log.Println("Created User!")
		redirectToHome(w, r)
	} else if err != nil {
		log.Println(err)
	} else {
		log.Println("User exists")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.PostFormValue("username")
	password := r.PostFormValue("password")

	db, err := OpenDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	id, pass, err := selectUser(db, login)
	if err == sql.ErrNoRows {
		log.Println("User doesn't exist")
		redirectToSignIn(w, r)
	} else if err != nil {
		log.Println(err)
	} else {
		if pass == password {
			log.Println("Logged in")
			currentUser = login //устанавливаем глобальное значение юзера
			currentUserId = id  //устанавливаем глобальное значение id юзера
			redirectToHome(w, r)
		} else {
			w.Write([]byte("Invalid password"))
		}
	}

}

func favoriteGet(w http.ResponseWriter, r *http.Request) {
	db, err := OpenDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	Books, err := selectFavoriteBook(db, currentUser, currentUserId)
	if err != nil {
		log.Println(err)
	}

	if len(Books) == 0 {
		err = Parse(w, "ui/html/favorite_without_books.html", currentUser)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = Parse(w, "ui/html/favorite_registered.html", Books[0])
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", 500)
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func favoriteUnSignedGet(w http.ResponseWriter, r *http.Request) {
	err := ParseNil(w, "ui/html/favorite_unregistered.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

// выход пользователя
func Exit(w http.ResponseWriter, r *http.Request) {
	currentUser = ""
	currentUserId = 0
	redirectToUnregistered(w, r)
}
