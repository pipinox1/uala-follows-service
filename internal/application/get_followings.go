package application

import (
	"context"
	"uala-followers-service/internal/domain"
)

type GetFollowingsCommand struct {
	UserID string
}

type GetFollowingsResponse struct {
	Followings []string `json:"followings"`
}

type GetFollowings struct {
	followRepository domain.FollowRepository
}

func NewGetFollowings(followRepository domain.FollowRepository) *GetFollowings {
	return &GetFollowings{
		followRepository: followRepository,
	}
}

func (g *GetFollowings) Exec(ctx context.Context, cmd *GetFollowingsCommand) (*GetFollowingsResponse, error) {
	followings, err := g.followRepository.FindFollowing(ctx, cmd.UserID)
	if err != nil {
		return nil, err
	}
	return &GetFollowingsResponse{
		Followings: followings,
	}, nil
}
