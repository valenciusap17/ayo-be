package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)


func NewDBConnection(user, password, dbname, port, host string) *sqlx.DB {
    connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", 
							user, password, dbname, port, host)
	
	fmt.Println("Connecting to db")
	db := sqlx.MustConnect("postgres", connectionString)
	
	if err := db.Ping(); err != nil {
		fmt.Println("Error while trying to connect to db")
	}
	fmt.Println("Successfully connected to db")
    return db
}