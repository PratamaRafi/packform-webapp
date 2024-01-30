/*
run this unit test with
go test ./controllers/Utest_customer_order.go
*/

package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"packform-webapp/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// Set up a test PostgreSQL database connection
	dsn := "user=test dbname=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Perform database migrations here if needed
	db.AutoMigrate(&models.Orders{}, &models.Customer{}, &models.Companies{}, &models.OrderItem{}, &models.Deliveries{})

	return db, nil
}

func TestGetAllOrderDetailsHandler(t *testing.T) {
	// Set up a test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}

	// Create an instance of OrderController with the test database
	orderController := NewOrderController(db)

	// Create a Gin router
	router := gin.Default()

	// Define the route for testing
	router.GET("/allOrderDetails", orderController.GetAllOrderDetailsHandler)

	// Perform a test HTTP request using httptest
	req, err := http.NewRequest("GET", "/allOrderDetails", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
