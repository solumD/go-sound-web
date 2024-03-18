package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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
