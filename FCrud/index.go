package main

import (
	_ "FCrud/mysql"
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

type Tabledata struct {
	ID        int
	Nama      string
	Deskripsi string
	Status    int
}

func konek() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "demo_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var klise = template.Must(template.ParseGlob("laman/*"))

func TList(w http.ResponseWriter, r *http.Request) {
	db := konek()
	pilDB, err := db.Query("SELECT * FROM demo_tbl_01 ORDER BY tbl1_id DESC")

	if err != nil {
		panic(err.Error())
	}
	tbl := Tabledata{}
	tbls := []Tabledata{}
	for pilDB.Next() {
		var tbl1_id int
		var tbl1_name, tbl1_desc string
		var tbl1_status int
		err = pilDB.Scan(&tbl1_id, &tbl1_name, &tbl1_desc, &tbl1_status)
		if err != nil {
			panic(err.Error())
		}
		tbl.ID = tbl1_id
		tbl.Nama = tbl1_name
		tbl.Deskripsi = tbl1_desc
		tbl.Status = tbl1_status

		tbls = append(tbls, tbl)
	}
	klise.ExecuteTemplate(w, "TList", tbls)
	defer db.Close()
}

func TTampil(w http.ResponseWriter, r *http.Request) {
	db := konek()
	gID := r.URL.Query().Get("ID")
	pilDB, err := db.Query("SELECT * FROM demo_tbl_01 WHERE tbl1_id=?", gID)

	if err != nil {
		panic(err.Error())
	}
	tbl := Tabledata{}
	for pilDB.Next() {
		var tbl1_id int
		var tbl1_name, tbl1_desc string
		var tbl1_status int
		err = pilDB.Scan(&tbl1_id, &tbl1_name, &tbl1_desc, &tbl1_status)
		if err != nil {
			panic(err.Error())
		}
		tbl.ID = tbl1_id
		tbl.Nama = tbl1_name
		tbl.Deskripsi = tbl1_desc
		tbl.Status = tbl1_status
	}
	klise.ExecuteTemplate(w, "TTampil", tbl)
	defer db.Close()
}

func TTambah(w http.ResponseWriter, r *http.Request) {
	klise.ExecuteTemplate(w, "TTambah", nil)
}

func TUbah(w http.ResponseWriter, r *http.Request) {
	db := konek()
	gID := r.URL.Query().Get("ID")
	pilDB, err := db.Query("SELECT * FROM demo_tbl_01 WHERE tbl1_id=?", gID)

	if err != nil {
		panic(err.Error())
	}
	tbl := Tabledata{}
	for pilDB.Next() {
		var tbl1_id int
		var tbl1_name, tbl1_desc string
		var tbl1_status int
		err = pilDB.Scan(&tbl1_id, &tbl1_name, &tbl1_desc, &tbl1_status)
		if err != nil {
			panic(err.Error())
		}
		tbl.ID = tbl1_id
		tbl.Nama = tbl1_name
		tbl.Deskripsi = tbl1_desc
		tbl.Status = tbl1_status
	}
	klise.ExecuteTemplate(w, "TUbah", tbl)
	defer db.Close()
}

func Tambah(w http.ResponseWriter, r *http.Request) {
	db := konek()
	if r.Method == "POST" {
		Nama := r.FormValue("Nama")
		Deskripsi := r.FormValue("Deskripsi")
		Status := r.FormValue("Status")

		tForm, err := db.Prepare("INSERT INTO demo_tbl_01(tbl1_name, tbl1_desc, tbl1_status) VALUES(?,?,?)")

		if err != nil {
			panic(err.Error)
		}
		tForm.Exec(Nama, Deskripsi, Status)
		log.Println("INSERT: NAMA : " + Nama + " | DESKRIPSI : " + Deskripsi)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Ubah(w http.ResponseWriter, r *http.Request) {
	db := konek()
	if r.Method == "POST" {
		Nama := r.FormValue("Nama")
		Deskripsi := r.FormValue("Deskripsi")
		Status := r.FormValue("Status")
		ID := r.FormValue("ID")
		uForm, err := db.Prepare("UPDATE demo_tbl_01 SET tbl1_name=?, tbl1_desc=?, tbl1_status=?, tbl1_id=?")

		if err != nil {
			panic(err.Error)
		}
		uForm.Exec(Nama, Deskripsi, Status, ID)
		log.Println("UPDATE: NAMA : " + Nama + " | DESKRIPSI : " + Deskripsi)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Hapus(w http.ResponseWriter, r *http.Request) {
	db := konek()
	tbl := r.URL.Query().Get("ID")
	delForm, err := db.Prepare("DELETE FROM demo_tbl_01 WHERE tbl1_id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(tbl)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", TList)
	http.HandleFunc("/ttampil", TTampil)
	http.HandleFunc("/ttambah", TTambah)
	http.HandleFunc("/tubah", TUbah)
	http.HandleFunc("/tambah", Tambah)
	http.HandleFunc("/ubah", Ubah)
	http.HandleFunc("/hapus", Hapus)
	http.ListenAndServe(":8000", nil)
}
