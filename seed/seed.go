package seed

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"

	// "log"
	"packform-webapp/models"
)

func SeedCustomer(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var customers []models.Customer

	for _, record := range records {
		// convert to list of string
		creditCards := pq.StringArray{record[5]}
		// Parse unit value
		uintValue, err := strconv.ParseUint(record[4], 10, 64)
		if err != nil {
			continue
		}
		customer := models.Customer{
			UserID:     record[0],
			Login:      record[1],
			Password:   record[2],
			Name:       record[3],
			CompanyID:  uint(uintValue),
			CreditCard: creditCards,
		}
		customers = append(customers, customer)
	}

	return db.Create(&customers).Error
}

func SeedCompanies(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var companies []models.Companies

	for _, record := range records {
		uintValue, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			continue
		}
		company := models.Companies{
			ID:           uint(uintValue),
			Company_Name: record[1],
		}
		companies = append(companies, company)
	}

	return db.Create(&companies).Error
}

func SeedDelivery(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var deliveries []models.Deliveries

	for _, record := range records {
		uintValue1, err1 := strconv.ParseUint(record[0], 10, 64)
		uintValue2, err2 := strconv.ParseUint(record[1], 10, 64)
		uintValue3, err3 := strconv.ParseUint(record[2], 10, 64)

		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}
		delivery := models.Deliveries{
			ID:                uint(uintValue1),
			OrderItemID:       uint(uintValue2),
			DeliveredQuantity: uint(uintValue3),
		}
		deliveries = append(deliveries, delivery)
	}

	return db.Create(&deliveries).Error
}

func SeedOrderItem(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var items []models.OrderItem

	for _, record := range records {

		if record[2] == "" {
			continue
		}

		uintValue1, err1 := strconv.ParseUint(record[0], 10, 64)
		uintValue2, err2 := strconv.ParseUint(record[1], 10, 64)
		price, err := strconv.ParseFloat(record[2], 32)
		uintValue3, err3 := strconv.ParseUint(record[3], 10, 64)

		if err != nil || err1 != nil || err2 != nil || err3 != nil {
			fmt.Printf("Error parsing price value: %v\n", err1, err2, err3)
			continue
		}

		item := models.OrderItem{

			ID:       uint(uintValue1),
			OrderID:  uint(uintValue2),
			Price:    float32(price),
			Quantity: uint(uintValue3),
			Prouct:   record[4],
		}
		items = append(items, item)
	}

	return db.Create(&items).Error
}

func SeedOrder(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var orders []models.Orders

	for _, record := range records {
		uintValue, err1 := strconv.ParseUint(record[0], 10, 64)
		// convert string to time.Time format
		time, err2 := time.Parse(time.RFC3339, record[1])
		if err1 != nil || err2 != nil {
			continue
		}

		order := models.Orders{
			ID:         uint(uintValue),
			CreatedAt:  time,
			OrderName:  record[2],
			CustomerID: record[3],
		}
		orders = append(orders, order)
	}

	return db.Create(&orders).Error
}
