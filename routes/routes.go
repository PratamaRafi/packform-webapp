package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"packform-webapp/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// CORS middleware configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://4pt4n6d6-8080.asse.devtunnels.ms/", "http://localhost:8081"} //local for staging domain
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type"}

	// Use the CORS middleware
	r.Use(cors.New(config))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	orderController := controllers.NewOrderController(db)

	// ORDER API
	r.GET("/order", orderController.GetAllOrderDetailsHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
