CREATE TABLE cis_asset_class (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    definition TEXT
);

CREATE TABLE cis_asset_criticality (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    value INTEGER NOT NULL
);

CREATE TABLE device_type (
    id SERIAL PRIMARY KEY,
    asset_class_id INTEGER NOT NULL REFERENCES cis_asset_class(id),
    name TEXT NOT NULL,
    UNIQUE(asset_class_id, name)
);

CREATE TABLE asset_source (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE asset (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    asset_class_id INTEGER NOT NULL REFERENCES cis_asset_class(id),
    criticality_id INTEGER NOT NULL REFERENCES cis_asset_criticality(id),
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE asset_source_map (
    asset_id INTEGER REFERENCES asset(id) ON DELETE CASCADE,
    asset_source_id INTEGER REFERENCES asset_source(id) ON DELETE CASCADE,
    PRIMARY KEY (asset_id, asset_source_id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS asset_source_map;
DROP TABLE IF EXISTS asset;
DROP TABLE IF EXISTS asset_source;
DROP TABLE IF EXISTS device_type;
DROP TABLE IF EXISTS cis_asset_criticality;
DROP TABLE IF EXISTS cis_asset_class;
