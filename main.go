package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	for {
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatal("This Error", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Cannot connect to DB:", err)
		}

		if err == nil && db.Ping() == nil {
			defer db.Close()
			log.Println("Connected to DB")
			break
		}

		log.Println("DB not ready, retrying in 2s...")
		time.Sleep(2 * time.Second)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from go with Postgres")
	})

	log.Println("Server Running on :8080")
	http.ListenAndServe(":8080", nil)
}
