package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register handles user registration.
// @Summary Register a new user
// @Description Register a new user with a username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user  body      RegisterRequest  true  "User registration info"
// @Success 201   {object}  map[string]string
// @Failure 400   {string}  string "Invalid request body"
// @Failure 500   {string}  string "Failed to create user"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := services.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login handles user login.
// @Summary Log in a user
// @Description Log in a user with a username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user  body      LoginRequest  true  "User login info"
// @Success 200   {object}  map[string]string
// @Failure 400   {string}  string "Invalid request body"
// @Failure 401   {string}  string "Invalid credentials"
// @Failure 500   {string}  string "Failed to generate token"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := services.LoginUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
