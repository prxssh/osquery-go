-- name: GetVersion :one
SELECT
    os_version,
    osquery_version
FROM
    versions;

-- name: UpsertVersions :exec
INSERT INTO versions (
    os_version,
    osquery_version
) VALUES (
    $1, $2
) ON CONFLICT (os_version, osquery_version) DO UPDATE SET
    os_version = $1,
    osquery_version = $2,
    updated_at = NOW();
