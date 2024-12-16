package main

import (
	"database/sql"
	"fmt"
	"lab_db_sql_queries/internal/api/insert"
	"lab_db_sql_queries/internal/database"
	"log"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	db_url := "postgres://postgres:@localhost:5432/lab?sslmode=disable"
	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not open a db connection %v", err))
	}

	queries := database.New(conn)

	api := ApiConfig{DB: queries}

	err = insert.InsertIntoMedPersonal(api.DB)
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = insert.InsertIntoWorkPlace(api.DB)
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = insert.InsertIntoOperationTypes(api.DB)
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = insert.InsertIntoWorkActivity(api.DB)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
