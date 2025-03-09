package user

import (
	"github.com/nynrathod/uber-ride/internal/ride"
	"time"

	"gorm.io/gorm"
)

// User represents the users in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Phone     string         `gorm:"uniqueIndex;size:20;not null" json:"phone"`
	Password  string         `gorm:"not null" json:"-"`
	Rides     []ride.Ride    `gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
