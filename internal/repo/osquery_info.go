package repo

import (
	"context"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/models"
	utils "github.com/prxssh/osquery-go/pkg"
)

type IOsqueryInfo interface {
	Get(ctx context.Context) (models.OsqueryInfo, error)
	Upsert(
		ctx context.Context,
		params map[string]any,
	) (models.OsqueryInfo, error)
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
	params map[string]any,
) (models.OsqueryInfo, error) {
	qtx := models.New(r.client)
	return qtx.Upsert(ctx, mapOsqueryInfoParams(params))
}

func (r *OsqueryInfoRepo) Get(
	ctx context.Context,
) (models.OsqueryInfo, error) {
	qtx := models.New(r.client)

	return qtx.GetOsqueryInfo(ctx)
}

func mapOsqueryInfoParams(data map[string]any) models.UpsertParams {
	params := models.UpsertParams{}

	utils.MapInt32Field(data, "pid", &params.Pid)
	utils.MapInt32Field(data, "config_valid", &params.ConfigValid)
	utils.MapInt32Field(data, "start_time", &params.StartTime)
	utils.MapInt32Field(data, "watcher", &params.Watcher)
	utils.MapInt32Field(data, "platform_mask", &params.PlatformMask)

	utils.MapStringField(data, "uuid", &params.Uuid)
	utils.MapStringField(data, "instance_id", &params.InstanceID)
	utils.MapStringField(data, "version", &params.Version)
	utils.MapStringField(data, "config_hash", &params.ConfigHash)
	utils.MapStringField(data, "extensions", &params.Extensions)
	utils.MapStringField(data, "build_platform", &params.BuildPlatform)
	utils.MapStringField(data, "build_distro", &params.BuildDistro)

	return params
}
