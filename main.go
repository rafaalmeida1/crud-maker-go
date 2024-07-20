package main

import (
	"database/sql"
    "net/http"
	"fmt"
	"log"

    "github.com/gorilla/mux"
)

const (
    host     = "db"
    port     = 5432
    user     = "user"
    password = "password"
    dbname   = "mydb"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	router := mux.NewRouter()

	RegisterApiV1(db, router)

    log.Println("Server starting on port 8000")
    if err := http.ListenAndServe(":8000", router); err != nil {
        log.Fatal(err)
    }
}
