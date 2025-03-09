package user

import (
	"errors"
)

// Service defines the business logic interface.
type Service interface {
	Register(u *User) error
	Login(email, password string) (*User, error)
}

type service struct {
	repo UserRepository
}

// NewService creates a new User service.
func NewService(repo UserRepository) Service {
	return &service{repo: repo}
}

// Register creates a new user if one doesn't already exist.
func (s *service) Register(u *User) error {
	// Check if user already exists.
	existing, err := s.repo.FindByEmail(u.Email)
	if err == nil && existing.ID != 0 {
		return errors.New("user already exists")
	}
	return s.repo.Create(u)
}

// Login verifies user credentials.
func (s *service) Login(email, password string) (*User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	// Naively compare passwords; in production, use proper hashing.
	if u.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return u, nil
}
