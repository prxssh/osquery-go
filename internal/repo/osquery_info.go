package repo

import (
	"context"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/models"
)

type IOsqueryInfo interface {
	List(ctx context.Context) (*models.OsqueryInfo, error)
	Upsert(ctx context.Context, osVersion models.OsqueryInfo) error
}

type OsqueryInfoRepo struct {
	client *postgres.PostgresClient
}

func NewOsqueryInfoRepo(dbClient *postgres.PostgresClient) IOsqueryInfo {
	return &OsqueryInfoRepo{
		client: dbClient,
	}
}

func (r *OsqueryInfoRepo) Upsert(
	ctx context.Context,
	osVersion models.OsqueryInfo,
) error {
	return nil
}

func (r *OsqueryInfoRepo) List(
	ctx context.Context,
) (*models.OsqueryInfo, error) {
	return nil, nil
}
