package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "github.com/lib/pq"
)

const (
    host     = "your-db-host"
    port     = "your-db-port"
    user     = "your-db-user"
    password = "your-db-password"
    dbname   = "your-db-name"
)

var db *sql.DB

func init() {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    // BAD: the category might have SQL special characters in it
    query1 := "SELECT ITEM, PRICE FROM PRODUCT WHERE ITEM_CATEGORY='" + r.URL.Query().Get("category") + "' ORDER BY PRICE"
    rows, err := db.Query(query1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // process results
    // ...

    // GOOD: use parameters
    query2 := "SELECT ITEM, PRICE FROM PRODUCT WHERE ITEM_CATEGORY=$1 ORDER BY PRICE"
    rows, err = db.Query(query2, r.URL.Query().Get("category"))
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // process results
    // ...

    // Respond to the client or do any additional processing
}

func main() {
    http.HandleFunc("/search", searchHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
