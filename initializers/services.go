package initializers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/BryceWayne/tiktak/pubsubclient"
	"github.com/BryceWayne/tiktak/repositories"
	"github.com/BryceWayne/tiktak/services"
)

func InitializeServices(ctx context.Context, projectID string) (services.UserService, services.VideoService, services.CommentService, services.LikeService, error) {
	pubSubService, err := InitializePubSubService(ctx, projectID)
	if err != nil {
		log.Fatalf("ERROR: InitializePubSubService - %v\n", err)
		return nil, nil, nil, nil, err // Return nils and the error
	}

	db, err := InitializeFirestore(ctx, projectID)
	if err != nil {
		log.Fatalf("ERROR: InitializeFirestore - %v\n", err)
		return nil, nil, nil, nil, err // Return nils and the error
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, pubSubService)

	videoRepo := repositories.NewVideoRepository(db)
	videoService := services.NewVideoService(videoRepo, pubSubService)

	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo, pubSubService)

	likeRepo := repositories.NewLikeRepository(db)
	likeService := services.NewLikeService(likeRepo, pubSubService)

	log.Println("SUCCESS: InitializeServices")
	return userService, videoService, commentService, likeService, nil
}

func InitializePubSubService(ctx context.Context, projectID string) (*pubsubclient.PubSubService, error) {
	pubSubClient, err := pubsubclient.NewPubSubService(ctx, projectID)
	if err != nil {
		log.Fatalf("ERROR: InitializePubSubService - %v\n", err)
		return nil, err
	}

	log.Println("SUCCESS: InitializePubSubService")
	return pubSubClient, nil
}

func InitializeFirestore(ctx context.Context, projectID string) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("ERROR: InitializeFirestore - %v", err)
		return nil, err
	}

	log.Println("SUCCESS: InitializeFirestore")
	return client, nil
}
