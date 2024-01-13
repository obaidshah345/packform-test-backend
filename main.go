package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Order represents the database model
type Order struct {
	ID                   uint    `json:"id"`
	OrderName            string  `json:"order_name"`
	CustomerCompanyName  string  `json:"customer_company_name"`
	CustomerName         string  `json:"customer_name"`
	OrderDate            string  `json:"order_date"`
	DeliveredAmount      float64 `json:"delivered_amount"`
	TotalAmount          float64 `json:"total_amount"`
}

var db *gorm.DB

func init() {
	// Initialize the database connection
	initDB()
}

func initDB() {
	// Set database credentials
	dbUsername := "postgres"
	dbPassword := "admin"
	dbName := "testdb123"
	dbPort	:= "5432"
	dbSSLMode := "disable"

	// Construct the connection string
	dbURI := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s", dbUsername, dbPassword, dbName, dbPort, dbSSLMode)

	// Connect to the database
	var err error
	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// AutoMigrate compares the table schema with the schema provided in struct for validity check
	db.AutoMigrate(&Order{})
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route to get all orders
	router.GET("/api/orders", getOrders)

	// Run the server
	router.Run(":8081")
}

// getOrders retrieves all orders from the database
func getOrders(c *gin.Context) {
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
