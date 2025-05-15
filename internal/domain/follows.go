package domain

import (
	"context"
	"time"
)

type FollowRepository interface {
	Create(ctx context.Context, follow *Follow) error
	Delete(ctx context.Context, followerID, followedID string) error
	FindFollowers(ctx context.Context, userID string) ([]string, error)
	FindFollowing(ctx context.Context, userID string) ([]string, error)
}

type Follow struct {
	FollowerID string
	FollowedID string
	CreatedAt  time.Time
}

func CreateFollow(followerID, followedID string) (*Follow, error) {
	return &Follow{
		FollowerID: followerID,
		FollowedID: followedID,
		CreatedAt:  time.Now(),
	}, nil
}
