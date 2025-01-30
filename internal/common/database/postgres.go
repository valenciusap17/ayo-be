package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Conn struct {
	db *sqlx.DB
}

func NewDBConnection() *Conn {
    connectionString := "user=postgres password=postgres dbname=ayo port=5432 host=localhost sslmode=disable"
	fmt.Println("Connecting to db")
	db := sqlx.MustConnect("postgres", connectionString)
	
	if err := db.Ping(); err != nil {
		fmt.Println("Error while trying to connect to db")
	}
	fmt.Println("Successfully connected to db")
    return &Conn{db: db}
}