package repositories

import (
    "context"

    "github.com/BryceWayne/tiktak/models"
)

type LikeRepository interface {
    SaveLike(ctx context.Context, like *models.Like) error
    GetLikesByVideoID(ctx context.Context, videoID string) ([]models.Like, error)
    // ... other methods for like CRUD operations
}

type likeRepository struct {
    // db connection or any other data source
}

func NewLikeRepository( /* db connection or other data sources */ ) LikeRepository {
    return &likeRepository{ /* ... */ }
}

func (repo *likeRepository) SaveLike(ctx context.Context, like *models.Like) error {
    // Implementation for saving a like in the database
}

func (repo *likeRepository) GetLikesByVideoID(ctx context.Context, videoID string) ([]models.Like, error) {
    // Implementation for retrieving likes by video ID from the database
}

// ... other CRUD operation implementations
