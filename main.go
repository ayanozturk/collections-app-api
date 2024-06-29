package main

import (
	"collections/database"
	"collections/handlers"
	"collections/repositories"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "app:app@tcp(localhost:3306)/app"
	}
	dataSourceName := dsn
	database.InitDB(dataSourceName)

	collectionRepo := repositories.NewCollectionRepository()
	handlers.InitHandlers(collectionRepo)

	router.GET("/collections", handlers.GetCollections)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
