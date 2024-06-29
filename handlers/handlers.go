package handlers

import (
	"collections/models"
	"collections/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var collectionRepo repositories.CollectionRepository

func InitHandlers(repo repositories.CollectionRepository) {
	collectionRepo = repo
}

func GetCollections(c *gin.Context) {
	collections, err := collectionRepo.GetCollections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if collections == nil {
		collections = []models.Collection{}
	}

	c.JSON(http.StatusOK, collections)
}

func AddCollection(c *gin.Context) {
	var newCollection models.Collection
	if err := c.ShouldBindJSON(&newCollection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err := collectionRepo.AddCollection(&newCollection)
	if err != nil {
		log.Printf("Error in AddCollection handler: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, newCollection)
}
