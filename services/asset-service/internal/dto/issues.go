package dtos

import "time"

// Issue Category

type CreateIssueCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateIssueCategory struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type IssueCategory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Issue Type

type CreateIssueType struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateIssueType struct {
	CategoryID  *int    `json:"category_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type IssueType struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Issue Status

type CreateIssueStatus struct {
	Name string `json:"name"`
}

type UpdateIssueStatus struct {
	Name *string `json:"name"`
}

type IssueStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Issue Source

type CreateIssueSource struct {
	Name string `json:"name"`
}

type UpdateIssueSource struct {
	Name *string `json:"name"`
}

type IssueSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Asset Issue

type CreateAssetIssue struct {
	AssetID         int     `json:"asset_id"`
	IssueTypeID     int     `json:"issue_type_id"`
	StatusID        int     `json:"status_id"`
	IssueSourceID   int     `json:"issue_source_id"`
	ExternalIssueID string  `json:"external_issue_id"`
	Description     *string `json:"description"`
}

type UpdateAssetIssue struct {
	IssueTypeID     *int    `json:"issue_type_id"`
	StatusID        *int    `json:"status_id"`
	IssueSourceID   *int    `json:"issue_source_id"`
	ExternalIssueID *string `json:"external_issue_id"`
	Description     *string `json:"description"`
}

type AssetIssue struct {
	ID              int       `json:"id"`
	AssetID         int       `json:"asset_id"`
	IssueTypeID     int       `json:"issue_type_id"`
	StatusID        int       `json:"status_id"`
	IssueSourceID   int       `json:"issue_source_id"`
	ExternalIssueID string    `json:"external_issue_id"`
	Description     *string `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

