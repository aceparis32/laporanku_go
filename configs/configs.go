package configs

import (
	"log"
	"os"
	"project-go-dasar/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBinit create connection to database
func DBInit(dsn string) *gorm.DB {
	// db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf&parseTime=True&loc=Local")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	var models = []interface{}{
		&models.Company{},
		&models.Customer{},
		&models.Note{},
		&models.Item{},
		&models.Log{},
		&models.Sales{},
		&models.User{},
	}

	db.AutoMigrate(models...)
	return db
}

// Database connection
func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		DisableAutomaticPing:   false,
	})
}
