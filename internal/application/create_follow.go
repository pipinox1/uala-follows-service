package application

import (
	"context"
	"uala-followers-service/internal/domain"
	"uala-followers-service/libs/events"
)

type CreateFollowCommand struct {
	FollowerID string `json:"-"`
	FollowedID string `json:"followed_id"`
}

type CreateFollow struct {
	followRepository domain.FollowRepository
	eventPublisher   events.Publisher
}

type CreateFollowResponse struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}

func NewCreateFollow(followRepository domain.FollowRepository, publisher events.Publisher) *CreateFollow {
	return &CreateFollow{
		followRepository: followRepository,
		eventPublisher:   publisher,
	}
}

func (c *CreateFollow) Exec(ctx context.Context, cmd *CreateFollowCommand) (*CreateFollowResponse, error) {
	follow, err := domain.CreateFollow(cmd.FollowerID, cmd.FollowedID)
	if err != nil {
		return nil, err
	}

	err = c.followRepository.Create(ctx, follow)
	if err != nil {
		return nil, err
	}

	return &CreateFollowResponse{
		FollowerID: follow.FollowerID,
		FollowedID: follow.FollowedID,
	}, nil
}
