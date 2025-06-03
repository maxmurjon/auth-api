package handler

import (
	"context"
	"fmt"
	"net/http"
	"smartlogistics/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	var entity *models.CreateProduct
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the Product in the storage
	id, err := h.strg.Product().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Product: " + err.Error(),
		})
		return
	}

	// Get the Product by ID
	product, err := h.strg.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Product: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Product has been created",
		Data:    product,
	})
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var entity models.UpdateProduct
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Update the Product
	if _, err := h.strg.Product().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Product: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Product has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetProductsList(c *gin.Context) {
	// Retrieve the list of Products (offset and limit can be implemented later)
	resp, err := h.strg.Product().GetList(context.Background(), &models.GetListProductRequest{})
	fmt.Println(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Product list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProductsByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Product list: " + err.Error(),
		})
		return
	}
	// Get the Product by ID
	product, err := h.strg.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Id: strconv.Itoa(intId)})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Product not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    product,
	})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Product list: " + err.Error(),
		})
		return
	}

	// Delete the Product by ID
	deletedProduct, err := h.strg.Product().Delete(context.Background(), &models.ProductPrimaryKey{Id: strconv.Itoa(intId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Product: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Product has been deleted",
		Data:    deletedProduct,
	})
}
