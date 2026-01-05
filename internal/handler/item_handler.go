package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufal225/go-simple-login-crud-api/internal/service"
)

type ItemHandler struct {
	itemService service.ItemService
}

func (h *ItemHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")

	var req struct {
		Name  string `json:"name"`
		SKU   string `json:"sku"`
		Stock int    `json:"stock"`
		Price int    `json:"price"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	item, err := h.itemService.Create(userID, req.Name, req.SKU, req.Price, req.Stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")

	items, err := h.itemService.List(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) Update(c *gin.Context) {
	userID := c.GetString("user_id")
	itemID := c.GetString("id")

	var req struct {
		Name string `json:"name"`
		Price int `json:"price"`
		Stock int `json:"stock"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
		return
	}

	err := h.itemService.Update(userID, itemID, req.Name, req.Price, req.Stock)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message":"item updated"})
}

func (h *ItemHandler) Delete(c *gin.Context) {
	userID := c.GetString("user_id")
	itemID := c.GetString("id")

	err := h.itemService.Delete(userID, itemID) 
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"item deleted"})
}

func NewItemHandler(itemService service.ItemService) *ItemHandler {
	return &ItemHandler{itemService: itemService}
}