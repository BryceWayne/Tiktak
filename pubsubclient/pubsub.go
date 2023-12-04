package pubsubclient

import (
    "context"

    "cloud.google.com/go/pubsub"
)

type PubSubService struct {
    client *pubsub.Client
}

func NewPubSubService(ctx context.Context, projectID string) (*PubSubService, error) {
    client, err := pubsub.NewClient(ctx, projectID)
    if err != nil {
        return nil, err
    }
    return &PubSubService{client: client}, nil
}

func (p *PubSubService) PublishMessage(ctx context.Context, topicName string, data []byte) error {
    topic := p.client.Topic(topicName)
    _, err := topic.Publish(ctx, &pubsub.Message{
        Data: data,
    }).Get(ctx)
    return err
}

// Add more methods as needed for subscribing, etc.
