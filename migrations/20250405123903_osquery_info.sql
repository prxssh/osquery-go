-- +goose Up
CREATE TABLE osquery_info (
    id SERIAL,
    pid INTEGER,
    uuid TEXT,
    instance_id TEXT,
    version TEXT,
    config_hash TEXT,
    config_valid INTEGER,
    extensions TEXT,
    build_platform TEXT,
    build_distro TEXT,
    start_time INTEGER,
    watcher INTEGER,
    platform_mask INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE (instance_id, uuid)
);

-- +goose Down
DROP TABLE osquery_info;
