CREATE TABLE device_detail (
    asset_id INTEGER PRIMARY KEY REFERENCES asset(id) ON DELETE CASCADE,
    hostname TEXT NOT NULL,
    device_type_id INTEGER NOT NULL REFERENCES device_type(id),
    ip_address TEXT,
    mac_address TEXT,
    os_name TEXT,
    os_version TEXT,
    hardware_cpe TEXT,
    last_seen TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE network_detail (
    asset_id INTEGER PRIMARY KEY REFERENCES asset(id) ON DELETE CASCADE,
    management_ip TEXT,
    device_type_id INTEGER NOT NULL REFERENCES device_type(id),
    mac_address TEXT,
    firmware_version TEXT,
    model_number TEXT,
    serial_number TEXT,
    total_ports INTEGER,
    last_seen TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE software_detail (
    asset_id INTEGER PRIMARY KEY REFERENCES asset(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    os_name TEXT,
    os_version TEXT,
    version TEXT,
    vendor TEXT,
    software_cpe TEXT,
    last_seen TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE cis_issue_status (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE cis_issue_category (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE issue_type (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES cis_issue_category(id),
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE issue_source (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE asset_issue (
    id SERIAL PRIMARY KEY,
    asset_id INTEGER NOT NULL REFERENCES asset(id) ON DELETE CASCADE,
    issue_type_id INTEGER NOT NULL REFERENCES issue_type(id),
    status_id INTEGER NOT NULL REFERENCES cis_issue_status(id),
    issue_source_id INTEGER NOT NULL REFERENCES issue_source(id),
    external_issue_id TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    UNIQUE(asset_id, issue_source_id, external_issue_id)
);


CREATE TABLE asset_note (
    id SERIAL PRIMARY KEY,
    asset_id INTEGER NOT NULL REFERENCES asset(id) ON DELETE CASCADE,
    asset_issue_id INTEGER REFERENCES asset_issue(id) ON DELETE SET NULL,
    author_user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS asset_issue;
DROP TABLE IF EXISTS issue_source;
DROP TABLE IF EXISTS issue_type;
DROP TABLE IF EXISTS cis_issue_category;
DROP TABLE IF EXISTS cis_issue_status;
DROP TABLE IF EXISTS software_detail;
DROP TABLE IF EXISTS network_detail;
DROP TABLE IF EXISTS device_detail;
DROP TABLE IF EXISTS asset_note;
