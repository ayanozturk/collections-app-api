package repositories

import (
	"collections/database"
	"collections/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserRepository interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	user.Password = string(hashedPassword)

	query := "INSERT INTO users (email, password) VALUES (:email, :password)"
	_, err = database.DB.NamedExec(query, user)
	// check if it's a duplicate entry error

	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

func (r *userRepository) Login(email, password string) (*models.User, error) {
	var user models.User
	err := database.DB.Get(&user, "SELECT id, email, password FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
