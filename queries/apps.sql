-- name: ListApps :many
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
    $2;

-- name: UpsertApp :one
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
          last_opened_time, created_at, updated_at;

-- name: CountApplications :one
SELECT
    COUNT(id)
FROM
    apps;
