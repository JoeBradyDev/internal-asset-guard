package service

import (
	"context"
	"encoding/json"
	"fmt"

	"asset-service/internal/db"
	"asset-service/internal/dtos"
	"asset-service/proto"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListAssets implements the paginated retrieval of assets with their JSONB details
func (s *AssetServer) ListAssets(ctx context.Context, req *proto.ListAssetsRequest) (*proto.ListAssetsResponse, error) {
	// 1. Execute paginated query from sqlc
	rows, err := s.Queries.ListFullAssetsPaged(ctx, db.ListFullAssetsPagedParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list assets: %w", err)
	}

	var assets []*proto.AssetResponse
	for _, row := range rows {
		assetRes := &proto.AssetResponse{
			Id:            row.ID,
			Name:          row.Name,
			AssetClass:    row.AssetClass,
			CriticalityId: row.CriticalityID,
			Criticality:   row.Criticality,
			CreatedAt:     timestamppb.New(row.CreatedAt.Time),
		}

		// 2. Unmarshal polymorphic JSONB details into Proto structs
		if len(row.DeviceInfo) > 0 {
			var d proto.DeviceDetail
			if err := json.Unmarshal(row.DeviceInfo, &d); err == nil {
				assetRes.DeviceInfo = &d
			}
		}
		if len(row.NetworkInfo) > 0 {
			var n proto.NetworkDetail
			if err := json.Unmarshal(row.NetworkInfo, &n); err == nil {
				assetRes.NetworkInfo = &n
			}
		}
		if len(row.SoftwareInfo) > 0 {
			var sw proto.SoftwareDetail
			if err := json.Unmarshal(row.SoftwareInfo, &sw); err == nil {
				assetRes.SoftwareInfo = &sw
			}
		}

		assets = append(assets, assetRes)
	}

	return &proto.ListAssetsResponse{Assets: assets}, nil
}

// CreateAsset maps the request to dtos.CreateAsset before persisting
func (s *AssetServer) CreateAsset(ctx context.Context, req *proto.CreateAssetRequest) (*proto.AssetResponse, error) {
	// 1. Use DTO for validation/logic
	dto := dtos.CreateAsset{
		Name:          req.Name,
		AssetClassID:  int(req.AssetClassId),
		CriticalityID: int(req.CriticalityId),
	}

	// 2. Create Core Asset
	asset, err := s.Queries.CreateAsset(ctx, db.CreateAssetParams{
		Name:          dto.Name,
		AssetClassID:  int32(dto.AssetClassID),
		CriticalityID: int32(dto.CriticalityID),
	})
	if err != nil {
		return nil, err
	}

	// 3. Create Details based on request
	if req.DeviceInfo != nil {
		_, err = s.Queries.CreateDeviceDetail(ctx, db.CreateDeviceDetailParams{
			AssetID:      asset.ID,
			Hostname:     pgtype.Text{String: req.DeviceInfo.GetHostname(), Valid: req.DeviceInfo.Hostname != nil},
			DeviceTypeID: pgtype.Int4{Int32: req.DeviceInfo.GetDeviceTypeId(), Valid: req.DeviceInfo.DeviceTypeId != nil},
			IpAddress:    pgtype.Text{String: req.DeviceInfo.GetIpAddress(), Valid: req.DeviceInfo.IpAddress != nil},
			MacAddress:   pgtype.Text{String: req.DeviceInfo.GetMacAddress(), Valid: req.DeviceInfo.MacAddress != nil},
			OsName:       pgtype.Text{String: req.DeviceInfo.GetOsName(), Valid: req.DeviceInfo.OsName != nil},
			OsVersion:    pgtype.Text{String: req.DeviceInfo.GetOsVersion(), Valid: req.DeviceInfo.OsVersion != nil},
			HardwareCpe:  pgtype.Text{String: req.DeviceInfo.GetHardwareCpe(), Valid: req.DeviceInfo.HardwareCpe != nil},
		})
	}

	return &proto.AssetResponse{
		Id:        asset.ID,
		Name:      asset.Name,
		CreatedAt: timestamppb.New(asset.CreatedAt.Time),
	}, err
}

// GetAsset fetches a single hydrated asset
func (s *AssetServer) GetAsset(ctx context.Context, req *proto.GetAssetRequest) (*proto.AssetResponse, error) {
	row, err := s.Queries.GetFullAsset(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res := &proto.AssetResponse{
		Id:            row.ID,
		Name:          row.Name,
		AssetClass:    row.AssetClass,
		CriticalityId: row.CriticalityID,
		Criticality:   row.Criticality,
		CreatedAt:     timestamppb.New(row.CreatedAt.Time),
	}

	// Unmarshal details
	if len(row.DeviceInfo) > 0 {
		json.Unmarshal(row.DeviceInfo, &res.DeviceInfo)
	}

	return res, nil
}

// UpdateAsset uses the fetch-merge-save pattern with dtos.UpdateAsset
func (s *AssetServer) UpdateAsset(ctx context.Context, req *proto.UpdateAssetRequest) (*proto.AssetResponse, error) {
	current, err := s.Queries.GetAssetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// Merge into DTO
	acid := int(current.AssetClassID)
	critid := int(current.CriticalityID)
	dto := dtos.UpdateAsset{
		Name:          &current.Name,
		AssetClassID:  &acid,
		CriticalityID: &critid,
	}

	if req.Name != nil { dto.Name = req.Name }
	if req.AssetClassId != nil {
		val := int(*req.AssetClassId)
		dto.AssetClassID = &val
	}

	updated, err := s.Queries.UpdateAsset(ctx, db.UpdateAssetParams{
		ID:            current.ID,
		Name:          *dto.Name,
		AssetClassID:  int32(*dto.AssetClassID),
		CriticalityID: int32(*dto.CriticalityID),
	})
	if err != nil {
		return nil, err
	}

	return &proto.AssetResponse{
		Id:        updated.ID,
		Name:      updated.Name,
		CreatedAt: timestamppb.New(updated.CreatedAt.Time),
	}, nil
}

func (s *AssetServer) DeleteAsset(ctx context.Context, req *proto.DeleteAssetRequest) (*proto.Empty, error) {
	err := s.Queries.DeleteAsset(ctx, req.Id)
	return &proto.Empty{}, err
}
