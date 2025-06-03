package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxmurjon/auth-api/config"
	"github.com/maxmurjon/auth-api/models"
	"github.com/maxmurjon/auth-api/pkg/helper"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary      Register new user
// @Description  Create new user with username and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.CreateUser  true  "User registration data"
// @Success      201   {object}  models.User
// @Failure      400   {object}  models.DefaultError  "Invalid input data"
// @Failure      409   {object}  models.DefaultError  "User already exists"
// @Failure      500   {object}  models.DefaultError  "Internal server error"
// @Router       /register [post]
func (h *Handler) Register(c *gin.Context) {
	var createUser models.CreateUser

	// JSONni bind qilish
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Error parsing registration data: " + err.Error(),
		})
		return
	}

	// Parolni hash qilish
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error hashing password: " + err.Error(),
		})
		return
	}

	createUser.Password = string(hashedPassword)

	// Foydalanuvchini yaratish
	userId, err := h.strg.User().Create(context.Background(), &createUser)
	if err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_login_key" (SQLSTATE 23505)` {
			c.JSON(http.StatusConflict, models.DefaultError{
				Message: "User already exists, please login!",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error creating user: " + err.Error(),
		})
		return
	}

	// Foydalanuvchi ma'lumotlarini olish
	user, err := h.strg.User().GetByID(context.Background(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error fetching user information: " + err.Error(),
		})
		return
	}

	// Yaratilgan foydalanuvchini qaytarish
	c.JSON(http.StatusCreated, user)
}

// Login godoc
// @Summary      Login user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      models.Login  true  "User login data"
// @Success      200    {object}  models.LoginResponse
// @Failure      400    {object}  models.DefaultError  "Invalid input or user not found"
// @Failure      401    {object}  models.DefaultError  "Unauthorized, invalid credentials"
// @Failure      500    {object}  models.DefaultError  "Internal server error"
// @Router       /login [post]
func (h *Handler) Login(c *gin.Context) {
	var login models.Login

	// JSONni bind qilish
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "Error parsing login data: " + err.Error()})
		return
	}

	// Foydalanuvchini telefon raqami bo'yicha olish
	resp, err := h.strg.User().GetByUserName(context.Background(), login.UserName)
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(http.StatusBadRequest, models.DefaultError{Message: "User not found, please register first"})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching user data: " + err.Error()})
		return
	}

	// Parollarni taqqoslash
	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.DefaultError{Message: "Invalid credentials"})
		return
	}

	// JWT token yaratish
	data := map[string]interface{}{
		"user_id":  resp.Id,
		"user_name": resp.UserName,
	}
	
	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.SekretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error generating JWT token: " + err.Error()})
		return
	}
	

	// JWT token va foydalanuvchi ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, models.LoginResponse{Token: token, UserData: resp})
}
