-- name: GetOsqueryInfo :one
SELECT 
    id,
    pid,
    uuid,
    instance_id,
    version,
    config_hash,
    config_valid,
    extensions,
    build_platform,
    build_distro,
    start_time,
    watcher,
    platform_mask,
    created_at,
    updated_at
FROM
    osquery_info;

-- name: Upsert :one
INSERT INTO osquery_info (
    pid,
    uuid,
    instance_id,
    version,
    config_hash,
    config_valid,
    extensions,
    build_platform,
    build_distro,
    start_time,
    watcher,
    platform_mask
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) ON CONFLICT (uuid) DO UPDATE SET
    pid = $1,
    version = $4,
    config_hash = $5,
    config_valid = $6,
    extensions = $7,
    build_platform = $8,
    build_distro = $9,
    start_time = $10,
    watcher = $11,
    platform_mask = $12,
    updated_at = NOW()
RETURNING id,
          pid,
          uuid,
          instance_id,
          version,
          config_hash,
          config_valid,
          extensions,
          build_platform,
          build_distro,
          start_time,
          watcher,
          platform_mask,
          created_at,
          updated_at;
