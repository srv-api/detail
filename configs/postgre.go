package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/plutov/paypal/v4"
	"github.com/srv-api/detail/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string for PostgreSQL
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	// Open connection to the database
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrate the schema
	db.AutoMigrate(
		&entity.Pin{},
		&entity.Subscribe{},
		&entity.Package{},
		&entity.CreditCard{},
		&entity.Merk{},
		&entity.Type{},
		&entity.Category{},
		&entity.Tax{},
		&entity.Discount{},
		&entity.Permission{},
		&entity.Table{},
		&entity.UserDetail{},
		&entity.Role{},
		&entity.RoleUser{},
		&entity.RoleUserPermission{},
		&entity.Printer{},
		&entity.PrinterAutoIncrement{},
		&entity.PaymentMethod{},
		&entity.UploadedPayment{},
		&entity.ContentSetting{},
		&entity.Message{},
		&entity.Order{},
		&entity.Reservation{},
		&entity.DeleteAccount{},
		&entity.FCMToken{},
		&entity.TopUp{},
		&entity.Qris{},
	)

	return db
}

func InitialMigration(db *gorm.DB) {
	// Get the underlying sql.DB instance from the gorm.DB instance.
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close() // Close the database connection
}

func InitApp() *paypal.Client {
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	secret := os.Getenv("PAYPAL_SECRET")
	paypalClient, err := paypal.NewClient(clientID, secret, paypal.APIBaseLive)
	if err != nil {
		log.Fatal("Failed to create PayPal client:", err)
	}

	_, err = paypalClient.GetAccessToken(context.Background())
	if err != nil {
		log.Fatal("Failed to get PayPal token:", err)
	}

	return paypalClient
}
