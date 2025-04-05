package repo

import (
	"context"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/models"
	utils "github.com/prxssh/osquery-go/pkg"
)

type IOsVersion interface {
	Get(ctx context.Context) (models.OsVersion, error)
	Upsert(ctx context.Context, params map[string]any) (models.OsVersion, error)
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
	params map[string]any,
) (models.OsVersion, error) {
	qtx := models.New(r.client)
	return qtx.UpsertOSDetails(ctx, mapOSDetailsParams(params))
}

func (r *OsVersionRepo) Get(ctx context.Context) (models.OsVersion, error) {
	qtx := models.New(r.client)

	return qtx.GetOSDetails(ctx)
}

func mapOSDetailsParams(data map[string]any) models.UpsertOSDetailsParams {
	params := models.UpsertOSDetailsParams{
		Name:    data["name"].(string),
		Version: data["version"].(string),
	}

	utils.MapStringField(data, "build", &params.Build)
	utils.MapStringField(data, "platform", &params.Platform)
	utils.MapStringField(data, "platform_like", &params.PlatformLike)
	utils.MapStringField(data, "codename", &params.Codename)
	utils.MapStringField(data, "arch", &params.Arch)
	utils.MapStringField(data, "extra", &params.Extra)
	utils.MapStringField(data, "mount_namespace_id", &params.MountNamespaceID)

	utils.MapInt32Field(data, "major", &params.Major)
	utils.MapInt32Field(data, "minor", &params.Minor)
	utils.MapInt32Field(data, "patch", &params.Patch)
	utils.MapInt32Field(data, "revision", &params.Revision)
	utils.MapInt32Field(data, "pid_with_namespace", &params.PidWithNamespace)

	utils.MapInt64Field(data, "install_date", &params.InstallDate)

	return params
}
