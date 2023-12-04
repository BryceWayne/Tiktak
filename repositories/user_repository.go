package repositories

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/BryceWayne/tiktak/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	// ... other methods for user CRUD operations
}

type userRepository struct {
	db *firestore.Client
}

func NewUserRepository(db *firestore.Client) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.Collection("users").Doc(user.ID).Set(ctx, user)
	log.Printf("SUCCESS: userRepository.CreateUser - %v", err)
	return err
}

func (repo *userRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	doc, err := repo.db.Collection("users").Doc(userID).Get(ctx)
	if err != nil {
		log.Printf("ERROR: userRepository.GetUserByID - %v", err)
		return nil, err
	}

	var user models.User
	err = doc.DataTo(&user)
	if err != nil {
		log.Printf("ERROR: userRepository.GetUserByID - %v", err)
		return nil, err
	}

	log.Printf("SUCCESS: userRepository.GetUserByID - %v", err)
	return &user, err
}

// ... other CRUD operation implementations
