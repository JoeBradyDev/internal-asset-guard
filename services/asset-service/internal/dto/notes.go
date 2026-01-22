package dtos

import "time"

// Asset Note

type CreateAssetNote struct {
	AssetID      int    `json:"asset_id"`
	AssetIssueID *int   `json:"asset_issue_id"`
	AuthorUserID int    `json:"author_user_id"`
	Content      string `json:"content"`
}

type UpdateAssetNote struct {
	AssetIssueID *int    `json:"asset_issue_id"`
	Content      *string `json:"content"`
}

type AssetNote struct {
	ID           int       `json:"id"`
	AssetID      int       `json:"asset_id"`
	AssetIssueID *int      `json:"asset_issue_id"`
	AuthorUserID int       `json:"author_user_id"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
}

// Asset Source Map

type AssetSourceMap struct {
	AssetID       int `json:"asset_id"`
	AssetSourceID int `json:"asset_source_id"`
}
