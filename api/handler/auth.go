package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxmurjon/auth-api/models"
	helper "github.com/maxmurjon/auth-api/pkg"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(c *gin.Context) {
	var input *models.Login
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Password hash
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    
	id,err:=h.strg.User().Create(c, &models.CreateUser{
		UserName: input.UserName,
		Password: string(hashedPassword),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

    user,err:=h.strg.User().GetByID(c, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
        return
    }

	
    c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been Registered",
		Data:    user,
	})
}

func (h *Handler) Login(c *gin.Context) {
    var input *models.Login
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user,err:=h.strg.User().GetByUserName(c, input.UserName)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }
    // Generate JWT
    token, err := helper.GenerateJWT(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
