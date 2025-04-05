-- +goose Up
CREATE TABLE apps (
    id SERIAL,
    name TEXT,
    path TEXT,
    bundle_executable TEXT,
    bundle_identifier TEXT,
    bundle_name TEXT,
    bundle_short_version TEXT,
    bundle_version TEXT,
    bundle_package_type TEXT,
    environment TEXT,
    element TEXT,
    compiler TEXT,
    development_region TEXT,
    display_name TEXT,
    info_string TEXT,
    minimum_system_version TEXT,
    category TEXT,
    applescript_enabled TEXT,
    copyright TEXT,
    last_opened_time DOUBLE PRECISION,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE (bundle_identifier, bundle_version)
);

-- +goose Down
DROP TABLE apps;
