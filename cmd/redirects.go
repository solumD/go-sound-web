package main

import (
	"net/http"
)

func redirectToSignIn(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, `http://localhost:8080/signin`, http.StatusSeeOther)
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, `http://localhost:8080/login`, http.StatusSeeOther)
}

func redirectToHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, `http://localhost:8080/home`, http.StatusSeeOther)
}

func redirectToUnregistered(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, `http://localhost:8080/`, http.StatusSeeOther)
}
