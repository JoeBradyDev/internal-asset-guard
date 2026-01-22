-- --- CIS ASSET CLASS ---

-- name: CreateAssetClass :one
INSERT INTO cis_asset_class (name, definition)
VALUES ($1, $2)
RETURNING *;

-- name: GetAssetClassByID :one
SELECT * FROM cis_asset_class
WHERE id = $1;

-- name: ListAssetClasses :many
SELECT * FROM cis_asset_class
ORDER BY name;

-- name: UpdateAssetClass :one
UPDATE cis_asset_class
SET name = $2, definition = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAssetClass :exec
DELETE FROM cis_asset_class
WHERE id = $1;

-- --- CIS ASSET CRITICALITY ---

-- name: CreateAssetCriticality :one
INSERT INTO cis_asset_criticality (name, value)
VALUES ($1, $2)
RETURNING *;

-- name: GetAssetCriticalityByID :one
SELECT * FROM cis_asset_criticality
WHERE id = $1;

-- name: ListAssetCriticalities :many
SELECT * FROM cis_asset_criticality
ORDER BY value DESC;

-- name: UpdateAssetCriticality :one
UPDATE cis_asset_criticality
SET name = $2, value = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAssetCriticality :exec
DELETE FROM cis_asset_criticality
WHERE id = $1;


-- --- DEVICE TYPE ---

-- name: CreateDeviceType :one
INSERT INTO device_type (asset_class_id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetDeviceTypeByID :one
SELECT * FROM device_type
WHERE id = $1;

-- name: ListDeviceTypesByClass :many
SELECT * FROM device_type
WHERE asset_class_id = $1
ORDER BY name;

-- name: UpdateDeviceType :one
UPDATE device_type
SET asset_class_id = $2, name = $3
WHERE id = $1
RETURNING *;

-- name: DeleteDeviceType :exec
DELETE FROM device_type
WHERE id = $1;


-- --- ISSUE CATEGORY ---

-- name: CreateIssueCategory :one
INSERT INTO cis_issue_category (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetIssueCategoryByID :one
SELECT * FROM cis_issue_category
WHERE id = $1;

-- name: ListIssueCategories :many
SELECT * FROM cis_issue_category
ORDER BY name;

-- name: UpdateIssueCategory :one
UPDATE cis_issue_category
SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteIssueCategory :exec
DELETE FROM cis_issue_category
WHERE id = $1;


-- --- ISSUE TYPE ---

-- name: CreateIssueType :one
INSERT INTO issue_type (category_id, name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetIssueTypeByID :one
SELECT * FROM issue_type
WHERE id = $1;

-- name: ListIssueTypesByCategory :many
SELECT * FROM issue_type
WHERE category_id = $1
ORDER BY name;

-- name: UpdateIssueType :one
UPDATE issue_type
SET category_id = $2, name = $3, description = $4
WHERE id = $1
RETURNING *;

-- name: DeleteIssueType :exec
DELETE FROM issue_type
WHERE id = $1;


-- --- ISSUE STATUS ---

-- name: CreateIssueStatus :one
INSERT INTO cis_issue_status (name)
VALUES ($1)
RETURNING *;

-- name: GetIssueStatusByID :one
SELECT * FROM cis_issue_status
WHERE id = $1;

-- name: ListIssueStatuses :many
SELECT * FROM cis_issue_status
ORDER BY id;

-- name: UpdateIssueStatus :one
UPDATE cis_issue_status
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteIssueStatus :exec
DELETE FROM cis_issue_status
WHERE id = $1;


-- --- ASSET SOURCE ---

-- name: CreateAssetSource :one
INSERT INTO asset_source (name)
VALUES ($1)
RETURNING *;

-- name: GetAssetSourceByID :one
SELECT * FROM asset_source
WHERE id = $1;

-- name: ListAssetSources :many
SELECT * FROM asset_source
ORDER BY name;

-- name: UpdateAssetSource :one
UPDATE asset_source
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAssetSource :exec
DELETE FROM asset_source
WHERE id = $1;


-- --- ISSUE SOURCE ---

-- name: CreateIssueSource :one
INSERT INTO issue_source (name)
VALUES ($1)
RETURNING *;

-- name: GetIssueSourceByID :one
SELECT * FROM issue_source
WHERE id = $1;

-- name: ListIssueSources :many
SELECT * FROM issue_source
ORDER BY name;

-- name: UpdateIssueSource :one
UPDATE issue_source
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteIssueSource :exec
DELETE FROM issue_source
WHERE id = $1;
