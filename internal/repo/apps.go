package repo

import (
	"context"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/models"
)

type IApps interface {
	List(ctx context.Context) ([]*models.App, error)
	Upsert(ctx context.Context, app models.App) error
}

type AppsRepo struct {
	client *postgres.PostgresClient
}

func NewAppsRepo(dbClient *postgres.PostgresClient) IApps {
	return &AppsRepo{
		client: dbClient,
	}
}

func (r *AppsRepo) Upsert(ctx context.Context, app models.App) error {
	return nil
}

func (r *AppsRepo) List(ctx context.Context) ([]*models.App, error) {
	return nil, nil
}
