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
       SELECT u.id
       FROM users u
       JOIN follows f ON f.follower_id = u.id
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
       SELECT u.id
       FROM users u
       JOIN follows f ON f.followed_id = u.id
       WHERE f.follower_id = $1`

	var following []string
	err := f.db.SelectContext(ctx, &following, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error finding following: %w", err)
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

func (f FollowRepository) Delete(ctx context.Context, followerID, followedID string) error {
	query := `
       DELETE FROM follows
       WHERE follower_id = $1 AND followed_id = $2`

	result, err := f.db.ExecContext(ctx, query, followerID, followedID)
	if err != nil {
		return fmt.Errorf("error deleting follow relationship: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (f FollowRepository) FindFollow(ctx context.Context, followerID, followedID string) (*domain.Follow, error) {
	query := `
       SELECT follower_id, followed_id, created_at
       FROM follows
       WHERE follower_id = $1 AND followed_id = $2`

	var followDB followDB
	err := f.db.GetContext(ctx, &followDB, query, followerID, followedID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding follow relationship: %w", err)
	}

	return followDB.toDomain(), nil
}

// CountFollowers cuenta el número de seguidores de un usuario
func (f FollowRepository) CountFollowers(ctx context.Context, userID string) (int, error) {
	query := `
       SELECT COUNT(*)
       FROM follows
       WHERE followed_id = $1`

	var count int
	err := f.db.GetContext(ctx, &count, query, userID)
	if err != nil {
		return 0, fmt.Errorf("error counting followers: %w", err)
	}

	return count, nil
}

// CountFollowing cuenta el número de usuarios que sigue un usuario
func (f FollowRepository) CountFollowing(ctx context.Context, userID string) (int, error) {
	query := `
       SELECT COUNT(*)
       FROM follows
       WHERE follower_id = $1`

	var count int
	err := f.db.GetContext(ctx, &count, query, userID)
	if err != nil {
		return 0, fmt.Errorf("error counting following: %w", err)
	}

	return count, nil
}

// FindFollowersWithPagination busca seguidores de un usuario con paginación
func (f FollowRepository) FindFollowersWithPagination(ctx context.Context, userID string, limit, offset int) ([]string, error) {
	query := `
       SELECT u.id
       FROM users u
       JOIN follows f ON f.follower_id = u.id
       WHERE f.followed_id = $1
       ORDER BY f.created_at DESC
       LIMIT $2 OFFSET $3`

	var followers []string
	err := f.db.SelectContext(ctx, &followers, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error finding followers with pagination: %w", err)
	}

	if len(followers) == 0 {
		return []string{}, nil
	}

	return followers, nil
}

// FindFollowingWithPagination busca usuarios seguidos por un usuario con paginación
func (f FollowRepository) FindFollowingWithPagination(ctx context.Context, userID string, limit, offset int) ([]string, error) {
	query := `
       SELECT u.id
       FROM users u
       JOIN follows f ON f.followed_id = u.id
       WHERE f.follower_id = $1
       ORDER BY f.created_at DESC
       LIMIT $2 OFFSET $3`

	var following []string
	err := f.db.SelectContext(ctx, &following, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error finding following with pagination: %w", err)
	}

	if len(following) == 0 {
		return []string{}, nil
	}

	return following, nil
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
