package repo

import (
	"context"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/models"
)

type IVersions interface {
	Upsert(ctx context.Context, params []map[string]any) error
	Get(ctx context.Context) (models.GetVersionRow, error)
}

type VersionsRepo struct {
	client *postgres.PostgresClient
}

func NewVersionsRepo(dbClient *postgres.PostgresClient) IVersions {
	return &VersionsRepo{
		client: dbClient,
	}
}

func (v *VersionsRepo) Upsert(
	ctx context.Context,
	params []map[string]any,
) error {
	qtx := models.New(v.client)
	return qtx.UpsertVersions(ctx, *parseParams(params))
}

func (v *VersionsRepo) Get(ctx context.Context) (models.GetVersionRow, error) {
	qtx := models.New(v.client)
	return qtx.GetVersion(ctx)
}

func parseParams(versions []map[string]any) *models.UpsertVersionsParams {
	res := &models.UpsertVersionsParams{}

	for _, versionMap := range versions {
		versionVal := versionMap["version_value"].(string)

		if versionMap["version_type"] == "OS" {
			res.OsVersion = versionVal
		} else {
			res.OsqueryVersion = versionVal
		}
	}

	return res
}
