package service

import (
	"context"
	"fmt"

	"asset-service/internal/db"
	"asset-service/internal/dto"
  "asset-service/internal/service/pb"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateIssue creates a new issue tied to an asset.
func (s *AssetServer) CreateIssue(ctx context.Context, req *pb.CreateIssueRequest) (*pb.IssueResponse, error) {
	// 1. Map to CreateAssetIssue DTO
	dto := dto.CreateAssetIssue{
		AssetID:         int(req.AssetId),
		IssueTypeID:     int(req.IssueTypeId),
		StatusID:        int(req.StatusId),
		IssueSourceID:   int(req.IssueSourceId),
		ExternalIssueID: req.ExternalIssueId,
		Description:     req.Description,
	}

	// 2. Persist to DB
	issue, err := s.Queries.CreateAssetIssue(ctx, db.CreateAssetIssueParams{
		AssetID:         int32(dto.AssetID),
		IssueTypeID:     int32(dto.IssueTypeID),
		StatusID:        int32(dto.StatusID),
		IssueSourceID:   int32(dto.IssueSourceID),
		ExternalIssueID: dto.ExternalIssueID,
		Description:     pgtype.Text{String: req.GetDescription(), Valid: req.Description != nil},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create issue: %w", err)
	}

	return s.mapIssueToProto(issue), nil
}

// GetIssue fetches a single issue by ID with joined metadata.
func (s *AssetServer) GetIssue(ctx context.Context, req *pb.GetIssueRequest) (*pb.IssueResponse, error) {
	issue, err := s.Queries.GetIssueByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("issue not found: %w", err)
	}

	// Mapping from the GetIssueByIDRow (which includes JOINed names)
	return &pb.IssueResponse{
		Id:              issue.ID,
		AssetId:         issue.AssetID,
		IssueTypeId:     issue.IssueTypeID,
		IssueType:       issue.TypeName,
		Category:        issue.CategoryName,
		StatusId:        issue.StatusID,
		Status:          issue.StatusName,
		IssueSourceId:   issue.IssueSourceID,
		IssueSource:     issue.SourceName,
		ExternalIssueId: issue.ExternalIssueID,
		Description:     &issue.Description.String,
		UpdatedAt:       timestamppb.New(issue.UpdatedAt.Time),
	}, nil
}

// ListIssuesByAsset retrieves all issues for a specific asset.
func (s *AssetServer) ListIssuesByAsset(ctx context.Context, req *pb.ListIssuesRequest) (*pb.ListIssuesResponse, error) {
	rows, err := s.Queries.GetIssuesByAssetID(ctx, req.AssetId)
	if err != nil {
		return nil, err
	}

	var issues []*pb.IssueResponse
	for _, row := range rows {
		issues = append(issues, &pb.IssueResponse{
			Id:              row.ID,
			AssetId:         row.AssetID,
			Status:          row.StatusName,
			IssueType:       row.TypeName,
			ExternalIssueId: row.ExternalIssueID,
			UpdatedAt:       timestamppb.New(row.UpdatedAt.Time),
		})
	}

	return &pb.ListIssuesResponse{Issues: issues}, nil
}

// UpdateIssue uses the fetch-merge-save pattern with dto.UpdateAssetIssue.
func (s *AssetServer) UpdateIssue(ctx context.Context, req *pb.UpdateIssueRequest) (*pb.IssueResponse, error) {
	// 1. FETCH current state
	current, err := s.Queries.GetIssueByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// 2. MERGE into DTO
	typeID := int(current.IssueTypeID)
	statusID := int(current.StatusID)
	sourceID := int(current.IssueSourceID)

	dto := dto.UpdateAssetIssue{
		IssueTypeID:     &typeID,
		StatusID:        &statusID,
		IssueSourceID:   &sourceID,
		ExternalIssueID: &current.ExternalIssueID,
	}
	if current.Description.Valid {
		dto.Description = &current.Description.String
	}

	// Apply updates from request
	if req.IssueTypeId != nil {
		v := int(*req.IssueTypeId)
		dto.IssueTypeID = &v
	}
	if req.StatusId != nil {
		v := int(*req.StatusId)
		dto.StatusID = &v
	}
	if req.Description != nil {
		dto.Description = req.Description
	}

	// 3. SAVE
	updated, err := s.Queries.UpdateAssetIssue(ctx, db.UpdateAssetIssueParams{
		ID:              current.ID,
		IssueTypeID:     int32(*dto.IssueTypeID),
		StatusID:        int32(*dto.StatusID),
		IssueSourceID:   int32(*dto.IssueSourceID),
		ExternalIssueID: *dto.ExternalIssueID,
		Description:     pgtype.Text{String: req.GetDescription(), Valid: req.Description != nil},
	})
	if err != nil {
		return nil, err
	}

	// Return a fresh hydrated view
	return s.GetIssue(ctx, &pb.GetIssueRequest{Id: updated.ID})
}

// DeleteIssue removes an issue.
func (s *AssetServer) DeleteIssue(ctx context.Context, req *pb.DeleteIssueRequest) (*pb.Empty, error) {
	// Note: Ensure you have DeleteAssetIssue in your sql queries
	err := s.Queries.DeleteAssetIssue(ctx, req.Id)
	return &pb.Empty{}, err
}

// Internal helper to map core db.AssetIssue to pb.IssueResponse
func (s *AssetServer) mapIssueToProto(i db.AssetIssue) *pb.IssueResponse {
	return &pb.IssueResponse{
		Id:              i.ID,
		AssetId:         i.AssetID,
		IssueTypeId:     i.IssueTypeID,
		StatusId:        i.StatusID,
		IssueSourceId:   i.IssueSourceID,
		ExternalIssueId: i.ExternalIssueID,
		Description:     &i.Description.String,
		UpdatedAt:       timestamppb.New(i.UpdatedAt.Time),
	}
}
