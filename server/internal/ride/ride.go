package ride

import (
	"time"

	"gorm.io/gorm"
)

// Ride represents the ride details
type Ride struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	DriverID  uint           `json:"driver_id"`
	Pickup    string         `gorm:"not null" json:"pickup"`
	Dropoff   string         `gorm:"not null" json:"dropoff"`
	Status    string         `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Price     float64        `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
