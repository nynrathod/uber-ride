package config

import (
	"fmt"
	"github.com/nynrathod/uber-ride/internal/driver"
	"github.com/nynrathod/uber-ride/internal/payment"
	"github.com/nynrathod/uber-ride/internal/ride"
	"github.com/nynrathod/uber-ride/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// DB is the global database instance
var DB *gorm.DB

// ConnectDB initializes PostgreSQL connection
func ConnectDB() {
	// Load from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore",
		EnvConfigs.DbHost,
		EnvConfigs.DbUser,
		EnvConfigs.DbPassword,
		EnvConfigs.DbName,
		EnvConfigs.DbPort,
	)
	fmt.Println("DSN:", dsn)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging for queries
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	// Set the global DB instance
	DB = db

	log.Println("✅ Connected to PostgreSQL successfully!")

	// Run Auto Migrations
	MigrateDB()
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// MigrateDB runs the auto-migrations
func MigrateDB() {
	err := DB.AutoMigrate(
		&user.User{},
		&driver.Driver{},
		&ride.Ride{},
		&payment.Payment{},
	)
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Database Migration Completed!")
}
