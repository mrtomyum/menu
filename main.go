package main

import (
	"log"
	"github.com/mrtomyum/menu/models"
	"github.com/mrtomyum/menu/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	DB_HOST = "tcp(nava.work:3306)"
	DB_NAME = "system"
	DB_USER = "root"
	DB_PASS = "mypass"
)

var dsn = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"

func main() {
	db, err := models.NewDB(dsn)
	if err != nil {
		log.Panic("NewDB() Error:", err)
	}

	c := controllers.Env{DB:db}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/menu", c.MenuAll).Methods("GET")
	log.Println("start Router GET MenuAll")
	r.HandleFunc("/api/v1/menu", c.MenuInsert).Methods("POST")
	log.Println("start Router POST MenuNew")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
