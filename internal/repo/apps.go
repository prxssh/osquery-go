package repo

import (
	"context"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/models"
	utils "github.com/prxssh/osquery-go/pkg"
)

type IApps interface {
	Count(ctx context.Context) (int64, error)
	Upsert(ctx context.Context, params map[string]any) error
	UpsertWithTx(ctx context.Context, params []map[string]any) error
	List(ctx context.Context, limit int32, offset int32) ([]models.App, error)
}

type AppsRepo struct {
	client *postgres.PostgresClient
}

func NewAppsRepo(dbClient *postgres.PostgresClient) IApps {
	return &AppsRepo{
		client: dbClient,
	}
}

func (r *AppsRepo) Count(ctx context.Context) (int64, error) {
	qtx := models.New(r.client)
	return qtx.CountApplications(ctx)
}

func (r *AppsRepo) Upsert(ctx context.Context, params map[string]any) error {
	qtx := models.New(r.client)
	_, err := qtx.UpsertApp(ctx, mapAppParams(params))
	return err
}

func (r *AppsRepo) UpsertWithTx(
	ctx context.Context,
	params []map[string]any,
) error {
	txn, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer txn.Rollback()

	qtx := models.New(txn)

	for _, param := range params {
		_, err := qtx.UpsertApp(ctx, mapAppParams(param))
		if err != nil {
			_ = txn.Rollback()
			return err
		}
	}

	return txn.Commit()
}

func (r *AppsRepo) List(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]models.App, error) {
	qtx := models.New(r.client)
	return qtx.ListApps(
		ctx,
		models.ListAppsParams{Limit: limit, Offset: offset},
	)
}

func mapAppParams(data map[string]any) models.UpsertAppParams {
	params := models.UpsertAppParams{}

	utils.MapStringField(data, "name", &params.Name)
	utils.MapStringField(data, "path", &params.Path)
	utils.MapStringField(data, "bundle_executable", &params.BundleExecutable)
	utils.MapStringField(data, "bundle_identifier", &params.BundleIdentifier)
	utils.MapStringField(data, "bundle_name", &params.BundleName)
	utils.MapStringField(
		data,
		"bundle_short_version",
		&params.BundleShortVersion,
	)
	utils.MapStringField(data, "bundle_version", &params.BundleVersion)
	utils.MapStringField(data, "bundle_package_type", &params.BundlePackageType)
	utils.MapStringField(data, "environment", &params.Environment)
	utils.MapStringField(data, "element", &params.Element)
	utils.MapStringField(data, "compiler", &params.Compiler)
	utils.MapStringField(data, "development_region", &params.DevelopmentRegion)
	utils.MapStringField(data, "display_name", &params.DisplayName)
	utils.MapStringField(data, "info_string", &params.InfoString)
	utils.MapStringField(
		data,
		"minimum_system_version",
		&params.MinimumSystemVersion,
	)
	utils.MapStringField(data, "category", &params.Category)
	utils.MapStringField(
		data,
		"applescript_enabled",
		&params.ApplescriptEnabled,
	)
	utils.MapStringField(data, "copyright", &params.Copyright)

	utils.MapFloat64Field(data, "last_opened_time", &params.LastOpenedTime)

	return params
}
