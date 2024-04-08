package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Coba sqlc")
	// buat koneksi ke postgresql dengan sqlc
	connStr := "postgresql://postgresql:12345@localhost:5432/sqlctest?sslmode=disable" // jadi instance variable untuk connection ke db psql"
	// postgresql://username db postgresql : password db postresql @localhost:5432 itu entry point dari db psqlnya  /sqlctest adalah nama dari databasenya
	// sslmode=disable berati nggak ada enkripsi antara aplikasi dan databasenya
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
