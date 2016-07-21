package controllers

import (
	"database/sql"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"github.com/mrtomyum/menu/models"
)

type Env struct {
	DB *sql.DB
}

func (e *Env) MenuAll(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	m := new(models.Menu)
	menus, err := m.All(e.DB)
	output, err := json.Marshal(menus)
	if err != nil {
		log.Println("Error json.Marshal:", err)
	}
	fmt.Fprintf(w, string(output))
}

func (e *Env) MenuInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(500), 500)
	}
	m := models.Menu{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		log.Println("Error decode.Decode(&m) >>", err)
	}
	err = m.Insert(e.DB)
	if err != nil {
		fmt.Println("Error Insert DB:", err)
	}
	output, _ := json.Marshal(&m)
	fmt.Fprintf(w, string(output))
}