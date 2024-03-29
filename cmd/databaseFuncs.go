package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	const connstr = "root:password@tcp(localhost:3306)/sound_web"
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("connected to database")
	return db, nil
}

// функция для проверки наличия пользователя в базе данных
func selectUser(db *sql.DB, login string) (string, error) {
	query := `select id, login, password from users where login = ?`
	row := db.QueryRow(query, login)
	var us User
	err := row.Scan(&us.id, &us.login, &us.password)
	if err != nil {
		return "", err
	}
	return us.password, nil
}

// функция для добавления пользователя в базу данных
func insertUser(db *sql.DB, login, password string) error {
	query := `insert into users(login, password) values(?, ?)`
	data := []any{login, password}
	res, err := db.Exec(query, data...)
	if err != nil {
		return nil
	}
	UserID, _ := res.LastInsertId()
	log.Printf("Last inserted id: %d", UserID)
	return nil
}
