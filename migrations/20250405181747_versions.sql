-- +goose Up
CREATE TABLE versions(
    id SERIAL,
    osquery_version TEXT NOT NULL,
    os_version TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE (osquery_version, os_version)
);
