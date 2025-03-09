package user

import (
	"gorm.io/gorm"
)

// UserRepository defines methods for interacting with the users table.
type UserRepository interface {
	Create(u *User) error
	FindByEmail(email string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	var u User
	err := r.db.Where("email = ?", email).First(&u).Error
	return &u, err
}
