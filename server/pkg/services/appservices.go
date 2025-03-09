package services

import (
	"github.com/nynrathod/uber-ride/internal/user"
)

type AppServices struct {
	UserService user.Service
	//AppsCollection *mongo.Collection
}
