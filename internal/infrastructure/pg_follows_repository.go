package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"uala-followers-service/internal/domain"
)

var _ domain.FollowRepository = (*FollowRepository)(nil)

type FollowRepository struct {
	db *sqlx.DB
}

func NewFollowRepository(db *sqlx.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

func (f FollowRepository) FindFollowers(ctx context.Context, userID string) ([]string, error) {
	query := `
       SELECT f.follower_id
       FROM follows f
       WHERE f.followed_id = $1`

	var followers []string
	err := f.db.SelectContext(ctx, &followers, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error finding followers: %w", err)
	}

	if len(followers) == 0 {
		return []string{}, nil
	}

	return followers, nil
}

func (f FollowRepository) FindFollowing(ctx context.Context, userID string) ([]string, error) {
	query := `
       SELECT f.followed_id
       FROM follows f
       WHERE f.follower_id = $1`

	var following []string
	err := f.db.SelectContext(ctx, &following, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrFollowNotFound
		}
		return nil, domain.ErrFollowInternalError
	}

	if len(following) == 0 {
		return []string{}, nil
	}

	return following, nil
}

func (f FollowRepository) Create(ctx context.Context, follow *domain.Follow) error {
	query := `
       INSERT INTO follows (follower_id, followed_id, created_at)
       VALUES (:follower_id, :followed_id, :created_at)`

	followDB := toDB(follow)
	_, err := f.db.NamedExecContext(ctx, query, followDB)
	if err != nil {
		return fmt.Errorf("error creating follow relationship: %w", err)
	}

	return nil
}

type followDB struct {
	FollowerID string    `db:"follower_id"`
	FollowedID string    `db:"followed_id"`
	CreatedAt  time.Time `db:"created_at"`
}

func toDB(follow *domain.Follow) followDB {
	return followDB{
		FollowerID: follow.FollowerID,
		FollowedID: follow.FollowedID,
		CreatedAt:  follow.CreatedAt,
	}
}

func (f followDB) toDomain() *domain.Follow {
	return &domain.Follow{
		FollowerID: f.FollowerID,
		FollowedID: f.FollowedID,
		CreatedAt:  f.CreatedAt,
	}
}
