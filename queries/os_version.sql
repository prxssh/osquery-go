-- name: GetOSDetails :one
SELECT 
    id,
    name,
    version,
    major,
    minor,
    patch,
    build,
    platform,
    platform_like,
    codename,
    arch,
    extra,
    install_date,
    revision,
    pid_with_namespace,
    mount_namespace_id,
    created_at,
    updated_at
FROM
    os_version;

--- name: UpsertOSDetails :one
INSERT INTO os_version (
    name,
    version,
    major,
    minor,
    patch,
    build,
    platform,
    platform_like,
    codename,
    arch,
    extra,
    install_date,
    revision,
    pid_with_namespace,
    mount_namespace_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) ON CONFLICT (platform, version) DO UPDATE SET
    name = $1,
    major = $3,
    minor = $4,
    patch = $5,
    build = $6,
    platform_like = $8,
    codename = $9,
    arch = $10,
    extra = $11,
    install_date = $12,
    revision = $13,
    pid_with_namespace = $14,
    mount_namespace_id = $15,
    updated_at = NOW()
RETURNING id, name, version, major, minor, patch, build, platform, platform_like, 
          codename, arch, extra, install_date, revision, pid_with_namespace, 
          mount_namespace_id, created_at, updated_at;
