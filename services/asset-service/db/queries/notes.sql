-- name: CreateAssetNote :one
INSERT INTO asset_note (
    asset_id,
    asset_issue_id,
    author_user_id,
    content
) VALUES (
    $1, -- asset_id
    $2, -- asset_issue_id (can be null)
    $3, -- author_user_id
    $4  -- content
)
RETURNING *;

-- name: GetNotesByAssetID :many
SELECT * FROM asset_note
WHERE asset_id = $1
ORDER BY created_at DESC;

-- name: UpdateAssetNote :one
UPDATE asset_note
SET
    asset_issue_id = $2,
    content = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAssetNote :exec
DELETE FROM asset_note WHERE id = $1;
