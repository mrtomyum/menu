package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Menu struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Path     string `json:"path"`
	Desc     string `json:"desc"`
}

type Menus []*Menu

func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("sql.Open() Error>>", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Panic("db.Ping() Error>>", err)
		return nil, err
	}
	log.Println("db = ", db)
	return db, nil //return db so in main can call defer db.Close()
}

func (m *Menu) All(db *sql.DB) ([]*Menu, error) {
	rows, err := db.Query(
		"SELECT id, parent_id, name, code, path  FROM menu")
	if err != nil {
		log.Println(">>> db.Query Error= ", err)
		return nil, err
	}
	defer rows.Close()

	var menus Menus
	for rows.Next() {
		m := new(Menu)
		err := rows.Scan(
			&m.ID,
			&m.ParentID,
			&m.Name,
			&m.Code,
			&m.Path,
		)
		if err != nil {
			log.Println(">>> rows.Scan() Error= ", err)
			return nil, err
		}
		menus = append(menus, m)
	}
	log.Println("Menu:", menus)
	return menus, nil
}

func (m *Menu) Insert(db *sql.DB) error {
	log.Println("Start m.New()")
	sql := "INSERT INTO menu (parent_id, name, code, path ) VALUES(?,?,?,?)"

	rs, err := db.Exec(sql,
		m.ParentID,
		m.Name,
		m.Code,
		m.Path,
	)
	if err != nil {
		log.Println(">>>Error cannot exec INSERT menu: >>>", err)
		return err
	}
	log.Println(rs)
	lastID, _ := rs.LastInsertId()

	err = db.QueryRow(
		"SELECT id, parent_id, name, code, path FROM menu WHERE id = ?",
		lastID,
	).Scan(
		&m.ID,
		&m.ParentID,
		&m.Name,
		&m.Code,
		&m.Path,
	)
	if err != nil {
		return err
	}
	log.Println("Success insert record:", m)
	return nil
}
