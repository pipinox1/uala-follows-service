package application

import (
	"context"
	"uala-followers-service/internal/domain"
)

type GetFollowersCommand struct {
	UserID string
}

type GetFollowersResponse struct {
	Followers []string `json:"followers"`
}

type GetFollowers struct {
	followRepository domain.FollowRepository
}

func NewGetFollowers(followRepository domain.FollowRepository) *GetFollowers {
	return &GetFollowers{
		followRepository: followRepository,
	}
}

func (g *GetFollowers) Exec(ctx context.Context, cmd *GetFollowersCommand) (*GetFollowersResponse, error) {
	followers, err := g.followRepository.FindFollowers(ctx, cmd.UserID)
	if err != nil {
		return nil, err
	}
	return &GetFollowersResponse{
		Followers: followers,
	}, nil
}
