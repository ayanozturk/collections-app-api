package handlers

import (
	"collections/models"
	"collections/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	collectionRepo repositories.CollectionRepository
	userRepo       repositories.UserRepository
)

func InitHandlers(collectionRepoParam repositories.CollectionRepository, userRepoParam repositories.UserRepository) {
	collectionRepo = collectionRepoParam
	userRepo = userRepoParam
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

func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err := userRepo.Register(&newUser)
	if err != nil {
		log.Printf("Error in RegisterUser handler: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	newUser.Password = ""
	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func LoginUser(c *gin.Context) {
	var loginRequest models.User
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	user, err := userRepo.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		log.Printf("Error in LoginUser handler: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}
