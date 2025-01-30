package main

import (
	"ayo/cmd/config"
	"ayo/internal/common/database"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// Connecting to postgresql db
	database.NewDBConnection()

	localConfig := config.Load()
	fmt.Printf("This is the result: %v", localConfig.Common.Postgres.User)
}