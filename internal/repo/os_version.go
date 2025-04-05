package repo

import (
	"context"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/models"
)

type IOsVersion interface {
	List(ctx context.Context) (*models.OsVersion, error)
	Upsert(ctx context.Context, osVersion models.OsVersion) error
}

type OsVersionRepo struct {
	client *postgres.PostgresClient
}

func NewOsVersionRepo(dbClient *postgres.PostgresClient) IOsVersion {
	return &OsVersionRepo{
		client: dbClient,
	}
}

func (r *OsVersionRepo) Upsert(
	ctx context.Context,
	osVersion models.OsVersion,
) error {
	return nil
}

func (r *OsVersionRepo) List(ctx context.Context) (*models.OsVersion, error) {
	return nil, nil
}
