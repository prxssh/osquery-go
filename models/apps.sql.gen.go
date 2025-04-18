// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: apps.sql

package models

import (
	"context"
	"database/sql"
)

const countApplications = `-- name: CountApplications :one
SELECT
    COUNT(id)
FROM
    apps
`

func (q *Queries) CountApplications(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countApplicationsStmt, countApplications)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listApps = `-- name: ListApps :many
SELECT
    id,
    name,
    path,
    bundle_executable,
    bundle_identifier,
    bundle_name,
    bundle_short_version,
    bundle_version,
    bundle_package_type,
    environment,
    element,
    compiler,
    development_region,
    display_name,
    info_string,
    minimum_system_version,
    category,
    applescript_enabled,
    copyright,
    last_opened_time,
    created_at,
    updated_at
FROM
    apps
ORDER BY 
    name
LIMIT 
    $1
OFFSET 
    $2
`

type ListAppsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListApps(ctx context.Context, arg ListAppsParams) ([]App, error) {
	rows, err := q.query(ctx, q.listAppsStmt, listApps, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []App{}
	for rows.Next() {
		var i App
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Path,
			&i.BundleExecutable,
			&i.BundleIdentifier,
			&i.BundleName,
			&i.BundleShortVersion,
			&i.BundleVersion,
			&i.BundlePackageType,
			&i.Environment,
			&i.Element,
			&i.Compiler,
			&i.DevelopmentRegion,
			&i.DisplayName,
			&i.InfoString,
			&i.MinimumSystemVersion,
			&i.Category,
			&i.ApplescriptEnabled,
			&i.Copyright,
			&i.LastOpenedTime,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertApp = `-- name: UpsertApp :one
INSERT INTO apps (
    name,
    path,
    bundle_executable,
    bundle_identifier,
    bundle_name,
    bundle_short_version,
    bundle_version,
    bundle_package_type,
    environment,
    element,
    compiler,
    development_region,
    display_name,
    info_string,
    minimum_system_version,
    category,
    applescript_enabled,
    copyright,
    last_opened_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
) ON CONFLICT (bundle_identifier, bundle_version) DO UPDATE SET
    name = $1,
    path = $2,
    bundle_executable = $3,
    bundle_name = $5,
    bundle_short_version = $6,
    bundle_package_type = $8,
    environment = $9,
    element = $10,
    compiler = $11,
    development_region = $12,
    display_name = $13,
    info_string = $14,
    minimum_system_version = $15,
    category = $16,
    applescript_enabled = $17,
    copyright = $18,
    last_opened_time = $19,
    updated_at = NOW()
RETURNING id, name, path, bundle_executable, bundle_identifier, bundle_name, 
          bundle_short_version, bundle_version, bundle_package_type, environment, 
          element, compiler, development_region, display_name, info_string, 
          minimum_system_version, category, applescript_enabled, copyright, 
          last_opened_time, created_at, updated_at
`

type UpsertAppParams struct {
	Name                 sql.NullString  `json:"name"`
	Path                 sql.NullString  `json:"path"`
	BundleExecutable     sql.NullString  `json:"bundle_executable"`
	BundleIdentifier     sql.NullString  `json:"bundle_identifier"`
	BundleName           sql.NullString  `json:"bundle_name"`
	BundleShortVersion   sql.NullString  `json:"bundle_short_version"`
	BundleVersion        sql.NullString  `json:"bundle_version"`
	BundlePackageType    sql.NullString  `json:"bundle_package_type"`
	Environment          sql.NullString  `json:"environment"`
	Element              sql.NullString  `json:"element"`
	Compiler             sql.NullString  `json:"compiler"`
	DevelopmentRegion    sql.NullString  `json:"development_region"`
	DisplayName          sql.NullString  `json:"display_name"`
	InfoString           sql.NullString  `json:"info_string"`
	MinimumSystemVersion sql.NullString  `json:"minimum_system_version"`
	Category             sql.NullString  `json:"category"`
	ApplescriptEnabled   sql.NullString  `json:"applescript_enabled"`
	Copyright            sql.NullString  `json:"copyright"`
	LastOpenedTime       sql.NullFloat64 `json:"last_opened_time"`
}

func (q *Queries) UpsertApp(ctx context.Context, arg UpsertAppParams) (App, error) {
	row := q.queryRow(ctx, q.upsertAppStmt, upsertApp,
		arg.Name,
		arg.Path,
		arg.BundleExecutable,
		arg.BundleIdentifier,
		arg.BundleName,
		arg.BundleShortVersion,
		arg.BundleVersion,
		arg.BundlePackageType,
		arg.Environment,
		arg.Element,
		arg.Compiler,
		arg.DevelopmentRegion,
		arg.DisplayName,
		arg.InfoString,
		arg.MinimumSystemVersion,
		arg.Category,
		arg.ApplescriptEnabled,
		arg.Copyright,
		arg.LastOpenedTime,
	)
	var i App
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Path,
		&i.BundleExecutable,
		&i.BundleIdentifier,
		&i.BundleName,
		&i.BundleShortVersion,
		&i.BundleVersion,
		&i.BundlePackageType,
		&i.Environment,
		&i.Element,
		&i.Compiler,
		&i.DevelopmentRegion,
		&i.DisplayName,
		&i.InfoString,
		&i.MinimumSystemVersion,
		&i.Category,
		&i.ApplescriptEnabled,
		&i.Copyright,
		&i.LastOpenedTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
