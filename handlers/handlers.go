package handlers

import (
	"collections/models"
	"collections/repositories"
	"github.com/gin-gonic/gin"
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
