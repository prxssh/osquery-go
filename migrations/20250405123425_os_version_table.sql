-- +goose Up
CREATE TABLE os_version (
    id SERIAL,
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    major INTEGER,
    minor INTEGER,
    patch INTEGER,
    build TEXT,
    platform TEXT,
    platform_like TEXT,
    codename TEXT,
    arch TEXT,
    extra TEXT,
    install_date BIGINT,
    revision INTEGER,
    pid_with_namespace INTEGER,
    mount_namespace_id TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE (name, version)
);

-- +goose Down
DROP TABLE os_version;
