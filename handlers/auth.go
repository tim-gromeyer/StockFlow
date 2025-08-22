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
// @Success 201   {object}  RegisterResponse
// @Failure 400   {object}  ErrorResponse "Invalid request body"
// @Failure 500   {object}  ErrorResponse "Failed to create user"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request body"})
		return
	}

	user, token, err := services.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, RegisterResponse{
		Username:    user.Username,
		CashBalance: user.CashBalance,
		Token:       token,
	})
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
// @Success 200   {object}  LoginResponse
// @Failure 400   {object}  ErrorResponse "Invalid request body"
// @Failure 401   {object}  ErrorResponse "Invalid credentials"
// @Failure 500   {object}  ErrorResponse "Failed to generate token"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request body"})
		return
	}

	token, user, err := services.LoginUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token, Username: user.Username})
}
