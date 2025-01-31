package main

import (
	"ayo/cmd/config"
	"ayo/cmd/dependency"
	"ayo/internal/account"
	"ayo/router"
	"context"
	"fmt"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	localConfig := config.Load()
	
	db := dependency.NewPostgreSQL(localConfig.Common.Postgres)
	ctx := context.Background()

	// Initialize adapters
	userSvc := account.NewAccountService(ctx, db)

	// Initialize ports
	routerEngine := gin.Default()
	routerEngine.Use(cors.Default())

	router.NewUserEngine(routerEngine, &localConfig.Routes.User, *userSvc)

	if err := routerEngine.Run(":8080"); err != nil {
		fmt.Printf("failed to run server %v", err)
	}
	fmt.Println("Application starts listering")
}