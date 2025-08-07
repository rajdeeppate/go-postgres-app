package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	dbURL := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("This Error", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from go with Postgres")
	})

	log.Println("Server Running on :8080")
	http.ListenAndServe(":8080", nil)
}
