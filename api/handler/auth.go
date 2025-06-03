package handler

// import (
// 	"context"
// 	"net/http"
// 	"smartlogistics/config"
// 	"smartlogistics/models"
// 	"smartlogistics/pkg/helper/helper"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// func (h *Handler) Register(c *gin.Context) {
// 	var createUser models.CreateUser

// 	// JSONni bind qilish
// 	err := c.ShouldBindJSON(&createUser)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.DefaultError{
// 			Message: "Error parsing registration data: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Parolni hash qilish
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{
// 			Message: "Error hashing password: " + err.Error(),
// 		})
// 		return
// 	}

// 	createUser.Password = string(hashedPassword)

// 	// Foydalanuvchini yaratish
// 	userId, err := h.strg.User().Create(context.Background(), &createUser)
// 	if err != nil {
// 		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_login_key" (SQLSTATE 23505)` {
// 			c.JSON(http.StatusConflict, models.DefaultError{
// 				Message: "User already exists, please login!",
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{
// 			Message: "Error creating user: " + err.Error(),
// 		})
// 		return
// 	}

// 	role, err := h.strg.Role().GetByName(context.Background(), &models.PrimaryKeyUUID{ID: "customer"})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{
// 			Message: "Error creating user: " + err.Error(),
// 		})
// 		return
// 	}

// 	_, err = h.strg.UserRole().Create(context.Background(), &models.UserRole{UserID: userId.Id, RoleID: role.Id})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{
// 			Message: "Error creating user: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Foydalanuvchi ma'lumotlarini olish
// 	user, err := h.strg.User().GetByID(context.Background(), userId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{
// 			Message: "Error fetching user information: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Yaratilgan foydalanuvchini qaytarish
// 	c.JSON(http.StatusCreated, user)
// }

// func (h *Handler) Login(c *gin.Context) {
// 	var login models.Login

// 	// JSONni bind qilish
// 	err := c.ShouldBindJSON(&login)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "Error parsing login data: " + err.Error()})
// 		return
// 	}

// 	// Foydalanuvchini telefon raqami bo'yicha olish
// 	resp, err := h.strg.User().GetByPhone(context.Background(), &login)
// 	if err != nil {
// 		if err.Error() == "no rows in result set" {
// 			c.JSON(http.StatusBadRequest, models.DefaultError{Message: "User not found, please register first"})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching user data: " + err.Error()})
// 		return
// 	}

// 	// Parollarni taqqoslash
// 	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, models.DefaultError{Message: "Invalid credentials"})
// 		return
// 	}

// 	userRole, err := h.strg.UserRole().GetByID(context.Background(), &models.PrimaryKeyUUID{ID: resp.Id})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching role data: " + err.Error()})
// 		return
// 	}

// 	roleId, err := strconv.Atoi(userRole.RoleID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching role data: " + err.Error()})
// 		return
// 	}

// 	role, err := h.strg.Role().GetByID(context.Background(), &models.PrimaryKey{Id: roleId})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching role data: " + err.Error()})
// 		return
// 	}

// 	// JWT token yaratish
// 	data := map[string]interface{}{
// 		"id":           resp.Id,
// 		"first_name":   resp.FirstName,
// 		"last_name":    resp.LastName,
// 		"phone_number": resp.PhoneNumber,
// 		"image_url":    resp.ImageUrl,
// 		// "role_id":      resp.RoleId,
// 		"created_at": resp.CreatedAt,
// 		"updated_at": resp.UpdatedAt,
// 		"role":       role.Id,
// 	}

// 	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.SekretKey)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error generating JWT token: " + err.Error()})
// 		return
// 	}

// 	// JWT token va foydalanuvchi ma'lumotlarini qaytarish
// 	c.JSON(http.StatusOK, models.LoginResponse{Token: token, UserData: resp})
// }
