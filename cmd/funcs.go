package main

import (
	"database/sql"
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

/*type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		log.Println(err)
	}
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}*/
