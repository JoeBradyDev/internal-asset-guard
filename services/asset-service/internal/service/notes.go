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

// CreateNote adds a new note to an asset or a specific issue.
func (s *AssetServer) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.NoteResponse, error) {
	// 1. Map to DTO
	var issueIDPtr *int
	if req.AssetIssueId != nil {
		id := int(*req.AssetIssueId)
		issueIDPtr = &id
	}

	dto := dto.CreateAssetNote{
		AssetID:      int(req.AssetId),
		AssetIssueID: issueIDPtr,
		AuthorUserID: int(req.AuthorUserId),
		Content:      req.Content,
	}

	// 2. Persist to DB
	note, err := s.Queries.CreateAssetNote(ctx, db.CreateAssetNoteParams{
		AssetID:      int32(dto.AssetID),
		AssetIssueID: pgtype.Int4{Int32: int32(req.GetAssetIssueId()), Valid: req.AssetIssueId != nil},
		AuthorUserID: int32(dto.AuthorUserID),
		Content:      dto.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create note: %w", err)
	}

	return &pb.NoteResponse{
		Id:           note.ID,
		AssetId:      note.AssetID,
		AssetIssueId: &note.AssetIssueID.Int32,
		Content:      note.Content,
		CreatedAt:    timestamppb.New(note.CreatedAt.Time),
	}, nil
}

// ListNotes retrieves all notes for a specific asset, ordered by newest first.
func (s *AssetServer) ListNotes(ctx context.Context, req *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	rows, err := s.Queries.GetNotesByAssetID(ctx, req.AssetId)
	if err != nil {
		return nil, err
	}

	var notes []*pb.NoteResponse
	for _, row := range rows {
		notes = append(notes, &pb.NoteResponse{
			Id:           row.ID,
			AssetId:      row.AssetID,
			AssetIssueId: &row.AssetIssueID.Int32,
			Content:      row.Content,
			CreatedAt:    timestamppb.New(row.CreatedAt.Time),
		})
	}

	return &pb.ListNotesResponse{Notes: notes}, nil
}

// UpdateNote handles the partial update of a note's content or issue link.
func (s *AssetServer) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.NoteResponse, error) {
	// 1. FETCH current state (Requires a GetNoteByID query in notes.sql)
	current, err := s.Queries.GetNoteByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("note not found: %w", err)
	}

	// 2. MERGE using DTO
	var currentIssueID *int
	if current.AssetIssueID.Valid {
		val := int(current.AssetIssueID.Int32)
		currentIssueID = &val
	}

	dto := dto.UpdateAssetNote{
		AssetIssueID: currentIssueID,
		Content:      &current.Content,
	}

	if req.Content != nil {
		dto.Content = req.Content
	}
	if req.AssetIssueId != nil {
		val := int(*req.AssetIssueId)
		dto.AssetIssueID = &val
	}

	// 3. SAVE
	updated, err := s.Queries.UpdateAssetNote(ctx, db.UpdateAssetNoteParams{
		ID:      current.ID,
		Content: *dto.Content,
		AssetIssueID: pgtype.Int4{
			Int32: int32(req.GetAssetIssueId()),
			Valid: req.AssetIssueId != nil,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.NoteResponse{
		Id:           updated.ID,
		Content:      updated.Content,
		AssetIssueId: &updated.AssetIssueID.Int32,
		CreatedAt:    timestamppb.New(updated.CreatedAt.Time),
	}, nil
}

// DeleteNote removes a note by ID.
func (s *AssetServer) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.Empty, error) {
	err := s.Queries.DeleteAssetNote(ctx, req.Id)
	return &pb.Empty{}, err
}
