package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrEmptyFollowedID     = errors.New("empty_followed")
	ErrFollowNotFound      = errors.New("follow.not_found")
	ErrFollowInternalError = errors.New("follow.internal_error")
)

type FollowRepository interface {
	Create(ctx context.Context, follow *Follow) error
	FindFollowers(ctx context.Context, userID string) ([]string, error)
	FindFollowing(ctx context.Context, userID string) ([]string, error)
}

type Follow struct {
	FollowerID string
	FollowedID string
	CreatedAt  time.Time
}

func CreateFollow(followerID, followedID string) (*Follow, error) {
	if followedID == "" {
		return nil, ErrEmptyFollowedID
	}
	return &Follow{
		FollowerID: followerID,
		FollowedID: followedID,
		CreatedAt:  time.Now(),
	}, nil
}
