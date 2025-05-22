package application

import (
	"context"
	"uala-followers-service/internal/domain"
)

type CreateFollowCommand struct {
	FollowerID string `json:"-"`
	FollowedID string `json:"followed_id"`
}

type CreateFollow struct {
	followRepository domain.FollowRepository
}

type CreateFollowResponse struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}

func NewCreateFollow(followRepository domain.FollowRepository) *CreateFollow {
	return &CreateFollow{
		followRepository: followRepository,
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
