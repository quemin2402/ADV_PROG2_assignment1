package handler

import (
	"Assignment1/Project/order-service/internal/domain"
	"Assignment1/Project/order-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	uc usecase.OrderUsecase
}

func NewOrderHandler(u usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{uc: u}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var o domain.Order
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.CreateOrder(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, o)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	o, err := h.uc.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, o)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var o domain.Order
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	o.ID = id
	if err := h.uc.UpdateOrder(&o); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, o)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.uc.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
