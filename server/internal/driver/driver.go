package driver

import (
	"github.com/nynrathod/uber-ride/internal/ride"
	"time"

	"gorm.io/gorm"
)

// Driver represents drivers in the system
type Driver struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Phone     string         `gorm:"uniqueIndex;size:20;not null" json:"phone"`
	CarModel  string         `gorm:"size:100" json:"car_model"`
	License   string         `gorm:"size:100;not null" json:"license"`
	Rides     []ride.Ride    `gorm:"foreignKey:DriverID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
