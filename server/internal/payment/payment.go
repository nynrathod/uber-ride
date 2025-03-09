package payment

import (
	"time"

	"gorm.io/gorm"
)

// Payment represents the payments for rides
type Payment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	RideID    uint           `gorm:"not null" json:"ride_id"`
	Amount    float64        `gorm:"not null" json:"amount"`
	Status    string         `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
