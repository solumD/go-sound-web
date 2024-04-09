package main

type User struct {
	id       int
	login    string
	password string
}

type Book struct {
	id     int
	User   string
	Title  string
	Author string
	Genre  string
}
