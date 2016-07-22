package main

import (
	"log"
	"github.com/mrtomyum/menu/models"
	"github.com/mrtomyum/menu/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

//TODO: Move Config to JSON file and Create Config{} to handle DB const.
const (
	DB_HOST = "tcp(nava.work:3306)"
	//TODO: เมื่อรันจริงต้องเปลี่ยนเป็น Docker Network Bridge IP เช่น 172.17.0.3 เป็นต้น
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
	r.HandleFunc("/api/v1/menu/tree", c.MenuTree).Methods("GET")
	log.Println("start Router GET MenuTree")

	http.Handle("/", r)
	http.ListenAndServe(":8001", nil)
}
