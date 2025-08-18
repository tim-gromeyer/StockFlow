package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

func RegisterUser(username string, password string) (*models.User, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &models.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		CashBalance:  10000, // Starting cash balance
	}

	if err := database.DB.Create(user).Error; err != nil {
		return nil, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}

func LoginUser(username string, password string) (string, *models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil, err
	}

	return tokenString, &user, nil
}

func GetJWTKey() []byte {
	return jwtKey
}
