package handlers

import (
	"collections/models"
	"collections/repositories"
	"collections/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	invalidInputMessage = gin.H{"message": "invalid input"}
	internalServerError = gin.H{"message": "internal server error"}
	collectionRepo      repositories.CollectionRepository
	userRepo            repositories.UserRepository
)

func InitHandlers(collectionRepoParam repositories.CollectionRepository, userRepoParam repositories.UserRepository) {
	collectionRepo = collectionRepoParam
	userRepo = userRepoParam
}

func GetCollections(c *gin.Context) {
	collections, err := collectionRepo.GetCollections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, internalServerError)
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
		c.JSON(http.StatusBadRequest, invalidInputMessage)
		return
	}

	err := collectionRepo.AddCollection(&newCollection)
	if err != nil {
		log.Printf("Error in AddCollection handler: %v", err)
		c.JSON(http.StatusInternalServerError, internalServerError)
		return
	}

	c.JSON(http.StatusCreated, newCollection)
}

func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, invalidInputMessage)
		return
	}

	err := userRepo.Register(&newUser)
	if err != nil {
		log.Printf("Error in RegisterUser handler: %v", err)
		c.JSON(http.StatusInternalServerError, internalServerError)
		return
	}
	newUser.Password = ""
	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func LoginUser(c *gin.Context) {
	var loginRequest models.User
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, invalidInputMessage)
		return
	}

	user, err := userRepo.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		log.Printf("Error in LoginUser handler: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	tokenString, refreshTokenString, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "authorization header missing"})
			c.Abort()
			return
		}

		// remove "Bearer " prefix
		tokenString = tokenString[7:]

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Next()
	}
}
