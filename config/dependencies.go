package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"uala-followers-service/internal/domain"
	"uala-followers-service/internal/infrastructure"
)

type Dependencies struct {
	FollowRepository domain.FollowRepository
}

func BuildDependencies(config Config) (*Dependencies, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Database,
	)
	if !config.Postgres.UseSSL {
		url += "?sslmode=disable"
	}
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	followRepository := infrastructure.NewFollowRepository(db)
	return &Dependencies{
		FollowRepository: followRepository,
	}, nil
}
