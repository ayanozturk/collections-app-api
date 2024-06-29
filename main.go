package main

import (
	"collections/database"
	"collections/handlers"
	"collections/repositories"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Update with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "*" // Update with your frontend URL
		//},
		MaxAge: 12 * time.Hour,
	}))

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "app:app@tcp(localhost:3306)/app"
	}
	dataSourceName := dsn
	database.InitDB(dataSourceName)

	collectionRepo := repositories.NewCollectionRepository()
	handlers.InitHandlers(collectionRepo)

	router.GET("/collections", handlers.GetCollections)
	router.POST("/collections", handlers.AddCollection)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
