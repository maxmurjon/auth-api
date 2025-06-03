package handler

import (
	"context"
	"net/http"
	"smartlogistics/models"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with given data
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.CreateUser  true  "User to create"
// @Success      200   {object}  models.SuccessResponse
// @Failure      400   {object}  models.DefaultError
// @Failure      500   {object}  models.DefaultError
// @Router       /createuser [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var entity *models.CreateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	id, err := h.strg.User().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create user: " + err.Error(),
		})
		return
	}

	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been created",
		Data:    user,
	})
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Update user details by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.UpdateUser  true  "User data to update"
// @Success      200   {object}  models.SuccessResponse
// @Failure      400   {object}  models.DefaultError
// @Failure      500   {object}  models.DefaultError
// @Router       /updateuser [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var entity models.UpdateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Yaroqsiz so'rov tanasi: " + err.Error(),
		})
		return
	}

	if entity.Id == "" {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Foydalanuvchi ID-si talab qilinadi",
		})
		return
	}

	if _, err := h.strg.User().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Foydalanuvchini yangilashda xato: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Foydalanuvchi yangilandi",
		Data:    entity.Id,
	})
}

// GetUsersList godoc
// @Summary      Get list of users
// @Description  Get paginated list of users
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.GetListUserResponse
// @Failure      500  {object}  models.DefaultError
// @Router       /users [get]
func (h *Handler) GetUsersList(c *gin.Context) {
	resp, err := h.strg.User().GetList(context.Background(), &models.GetListUserRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve user list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUsersByIDHandler godoc
// @Summary      Get user by ID
// @Description  Get user details by ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      404  {object}  models.DefaultError
// @Router       /user/{id} [get]
func (h *Handler) GetUsersByIDHandler(c *gin.Context) {
	id := c.Param("id")

	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "User not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    user,
	})
}

// DeleteUser godoc
// @Summary      Delete user by ID
// @Description  Delete user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      500  {object}  models.DefaultError
// @Router       /deleteuser/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	deletedUser, err := h.strg.User().Delete(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been deleted",
		Data:    deletedUser,
	})
}
