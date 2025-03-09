package services

import (
	cfg "github.com/nynrathod/uber-ride/config"
	usr "github.com/nynrathod/uber-ride/internal/user"
)

// AppServiceInitializer initializes and bundles your application services.
type AppServiceInitializer struct{}

// NewAppServiceInitializer creates a new AppServiceInitializer.
func NewAppServiceInitializer() *AppServiceInitializer {
	return &AppServiceInitializer{}
}

func (initializer *AppServiceInitializer) InitializeAppServices() AppServices {
	db := cfg.GetDB()
	if db == nil {
		panic("Database connection is not initialized")
	}

	userRepo := usr.NewUserRepository(db)

	userService := usr.NewService(userRepo)

	return AppServices{
		UserService: userService,
	}
}
