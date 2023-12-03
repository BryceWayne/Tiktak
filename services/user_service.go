package services

import (
    "context"

    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/repositories"
)

type UserService interface {
    RegisterUser(ctx context.Context, user *models.User) error
    GetUser(ctx context.Context, id string) (*models.User, error)
    // ... other business logic methods
}

type userService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo}
}

func (s *userService) RegisterUser(ctx context.Context, user *models.User) error {
    // Business logic for registering a user
    // This may involve calling userRepo.CreateUser and additional logic

    return nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*models.User, error) {
    // Business logic for getting a user
    // This may involve calling userRepo.GetUserByID and additional logic

    return &models.User{}, nil
}

// ... implementation of other methods
