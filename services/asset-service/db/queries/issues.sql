-- name: CreateAssetIssue :one
INSERT INTO asset_issue (
    asset_id,
    issue_type_id,
    status_id,
    issue_source_id,
    external_issue_id,
    description
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetIssueByID :one
SELECT
    i.*, cat.name as category_name, it.name as type_name, stat.name as status_name, isrc.name as source_name
FROM asset_issue i
JOIN issue_type it ON i.issue_type_id = it.id
JOIN cis_issue_category cat ON it.category_id = cat.id
JOIN cis_issue_status stat ON i.status_id = stat.id
JOIN issue_source isrc ON i.issue_source_id = isrc.id
WHERE i.id = $1;

-- name: GetIssuesByAssetID :many
SELECT
    i.*, cat.name as category_name, it.name as type_name, stat.name as status_name, isrc.name as source_name
FROM asset_issue i
JOIN issue_type it ON i.issue_type_id = it.id
JOIN cis_issue_category cat ON it.category_id = cat.id
JOIN cis_issue_status stat ON i.status_id = stat.id
JOIN issue_source isrc ON i.issue_source_id = isrc.id
WHERE i.asset_id = $1;

-- name: UpdateAssetIssue :one
UPDATE asset_issue
SET
    issue_type_id = $2,
    status_id = $3,
    issue_source_id = $4,
    external_issue_id = $5,
    description = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAssetIssue :exec
DELETE FROM asset_issue
WHERE id = $1;

