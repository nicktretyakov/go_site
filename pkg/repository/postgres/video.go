package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"website.friendsonly.me/pkg/repository"
)

// videoRepoPostgres is a users repository for postgresql DB
type videoRepoPostgres struct {
	pool   *pgxpool.Pool
	logger *logrus.Logger
}

// NewVideoRepo returns new VideoRepo object via conn string.
func NewVideoRepo(ctx context.Context, pool *pgxpool.Pool, logger *logrus.Logger) repository.VideoRepo {
	return &videoRepoPostgres{
		pool:   pool,
		logger: logger,
	}
}

// GetAll return list of all video records
func (v *videoRepoPostgres) GetAll(ctx context.Context) error {
	return nil
}
