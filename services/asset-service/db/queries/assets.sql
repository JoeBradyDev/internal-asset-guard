--- 1. ASSET CORE ---

-- name: CreateAsset :one
INSERT INTO asset (name, asset_class_id, criticality_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAssetByID :one
SELECT * FROM asset WHERE id = $1;

-- name: UpdateAsset :one
UPDATE asset
SET name = $2, asset_class_id = $3, criticality_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAsset :exec
DELETE FROM asset WHERE id = $1;

--- 2. ASSET SOURCE MAPPING ---

-- name: AddAssetSourceMap :exec
INSERT INTO asset_source_map (asset_id, asset_source_id)
VALUES ($1, $2);

-- name: GetSourcesByAsset :many
SELECT s.* FROM asset_source s
JOIN asset_source_map asm ON s.id = asm.asset_source_id
WHERE asm.asset_id = $1;

-- name: RemoveAssetFromSource :exec
DELETE FROM asset_source_map
WHERE asset_id = $1 AND asset_source_id = $2;

--- 3. DETAILS (DEVICE, NETWORK, SOFTWARE) ---

-- name: CreateDeviceDetail :one
INSERT INTO device_detail (
    asset_id, hostname, device_type_id, ip_address, mac_address, os_name, os_version, hardware_cpe, last_seen
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetDeviceDetail :one
SELECT * FROM device_detail
WHERE asset_id = $1;

-- name: UpdateDeviceDetail :one
UPDATE device_detail
SET
    hostname = $2,
    device_type_id = $3,
    ip_address = $4,
    mac_address = $5,
    os_name = $6,
    os_version = $7,
    hardware_cpe = $8,
    last_seen = $9
WHERE asset_id = $1
RETURNING *;

-- name: DeleteDeviceDetail :exec
DELETE FROM device_detail WHERE asset_id = $1;

-- name: CreateNetworkDetail :one
INSERT INTO network_detail (
    asset_id, management_ip, device_type_id, mac_address, firmware_version, model_number, serial_number, total_ports, last_seen
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetNetworkDetail :one
SELECT * FROM network_detail
WHERE asset_id = $1;

-- name: UpdateNetworkDetail :one
UPDATE network_detail
SET
    management_ip = $2,
    device_type_id = $3,
    mac_address = $4,
    firmware_version = $5,
    model_number = $6,
    serial_number = $7,
    total_ports = $8,
    last_seen = $9
WHERE asset_id = $1
RETURNING *;

-- name: DeleteNetworkDetail :exec
DELETE FROM network_detail WHERE asset_id = $1;

-- name: CreateSoftwareDetail :one
INSERT INTO software_detail (asset_id, name, os_name, os_version, version, vendor, software_cpe, last_seen)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetSoftwareDetail :one
SELECT * FROM software_detail
WHERE asset_id = $1;

-- name: UpdateSoftwareDetail :one
UPDATE software_detail
SET
    name = $2,
    os_name = $3,
    os_version = $4,
    version = $5,
    vendor = $6,
    software_cpe = $7,
    last_seen = $8
WHERE asset_id = $1
RETURNING *;

-- name: DeleteSoftwareDetail :exec
DELETE FROM software_detail WHERE asset_id = $1;

--- 4. COMPLEX RETRIEVAL ---

-- name: GetFullAsset :one
SELECT
    a.id, a.name, a.created_at,
    cl.name AS asset_class,
    cr.id AS criticality_id,
    cr.name AS criticality,
    (SELECT row_to_json(dt) FROM (
        SELECT d.*, l.name as device_type
        FROM device_detail d
        JOIN device_type l ON d.device_type_id = l.id
        WHERE d.asset_id = a.id
    ) dt) AS device_info,
    (SELECT row_to_json(nt) FROM (
        SELECT n.*, l.name as device_type
        FROM network_detail n
        JOIN device_type l ON n.device_type_id = l.id
        WHERE n.asset_id = a.id
    ) nt) AS network_info,
    (SELECT row_to_json(s) FROM software_detail s WHERE s.asset_id = a.id) AS software_info
FROM asset a
JOIN cis_asset_class cl ON a.asset_class_id = cl.id
JOIN cis_asset_criticality cr ON a.criticality_id = cr.id
WHERE a.id = $1;

-- name: ListFullAssetsPaged :many
SELECT
    a.id, a.name, a.created_at,
    cl.name AS asset_class,
    cr.id AS criticality_id,
    cr.name AS criticality,
    (SELECT row_to_json(dt) FROM (
        SELECT d.*, l.name as device_type
        FROM device_detail d
        JOIN device_type l ON d.device_type_id = l.id
        WHERE d.asset_id = a.id
    ) dt) AS device_info,
    (SELECT row_to_json(nt) FROM (
        SELECT n.*, l.name as device_type
        FROM network_detail n
        JOIN device_type l ON n.device_type_id = l.id
        WHERE n.asset_id = a.id
    ) nt) AS network_info,
    (SELECT row_to_json(s) FROM software_detail s WHERE s.asset_id = a.id) AS software_info,
    (SELECT count(*) FROM asset_issue WHERE asset_id = a.id) as total_issues
FROM asset a
JOIN cis_asset_class cl ON a.asset_class_id = cl.id
JOIN cis_asset_criticality cr ON a.criticality_id = cr.id
WHERE ($3::int = 0 OR a.asset_class_id = $3)
  AND ($4::int = 0 OR a.criticality_id = $4)
ORDER BY a.id ASC
LIMIT $1 OFFSET $2;
