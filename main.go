package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create and use dev database in docker.
	db, err := sql.Open("mysql", "rfamro:rfamro@tcp(mysql-rfam-public.ebi.ac.uk:4497)/Rfam")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! %s\n Request: %s %s\n Host: %s\n", r.Host, r.URL.Path, r.Method, r.URL.Host)
	})

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}