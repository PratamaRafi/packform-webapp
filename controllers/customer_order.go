package controllers

import (
	"net/http"
	"packform-webapp/models" // Import your actual models package

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (c *OrderController) GetAllCustomerOrder() ([]models.CustomerOrder, error) {
	var customerOrder []models.CustomerOrder

	result := c.DB.
		Table("orders").
		Select("orders.order_name,companies.company_name, customers.name,orders.created_at, (order_items.price * deliveries.delivered_quantity), (order_items.price * order_items.quantity) as total_amount,order_items.price,deliveries.delivered_quantity").
		Joins("JOIN customers on customers.user_id = orders.customer_id").
		Joins("JOIN companies  on companies.id =  customers.company_id").
		Joins("JOIN order_items  on order_items.order_id = order_items.id").
		Joins("JOIN deliveries  on deliveries.order_item_id = order_items.id").
		Scan(&customerOrder)

	if result.Error != nil {
		return nil, result.Error
	}

	return customerOrder, nil
}

func (c *OrderController) GetAllOrderDetailsHandler(ctx *gin.Context) {
	allOrderDetails, err := c.GetAllCustomerOrder()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": allOrderDetails})
}
