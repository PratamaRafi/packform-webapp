package main

import (
	"log"
	"packform-webapp/config"
	"packform-webapp/docs"
	"packform-webapp/routes"
	"packform-webapp/seed"
	"packform-webapp/utils"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// for load godotenv
	// for env
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// tokenString := token.ExtractToken(c)
	// fmt.Println(tokenString)

	// bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDQ0MTkzOTQsInVzZXJfaWQiOjJ9.B1KTe0DN8FtvsTorqPcuYi6_GQ6Kvfe1Nby3OVjU3Y4"
	// splitToken := strings.Split(bearerToken, "Bearer ")
	// if len(splitToken) != 2 {
	// 	fmt.Println("jeng jeng")
	// }
	// fmt.Println(splitToken[1])
	// bearerToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDQ0MTkzOTQsInVzZXJfaWQiOjJ9.B1KTe0DN8FtvsTorqPcuYi6_GQ6Kvfe1Nby3OVjU3Y4"
	// if len(strings.Split(bearerToken, " ")) == 2 {
	// 	fmt.Println(strings.Split(bearerToken, " ")[1])
	// }
	// fmt.Println("jeng jeng")

	docs.SwaggerInfo.Title = "Swagger Packform tech test API"
	docs.SwaggerInfo.Description = "This is a Packform tech test."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// database connection
	db, err := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err = seed.SeedCustomer(db, "test_data/customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = seed.SeedCompanies(db, "test_data/customer_companies.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = seed.SeedDelivery(db, "test_data/deliveries.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = seed.SeedOrderItem(db, "test_data/order_items.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = seed.SeedOrder(db, "test_data/orders.csv")
	if err != nil {
		log.Fatal(err)
	}

	// router
	r := routes.SetupRouter(db)
	r.Run("127.0.0.1:8080")
}
