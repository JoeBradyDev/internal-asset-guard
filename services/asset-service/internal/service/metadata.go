package service

import (
	"context"
	"fmt"

	"asset-service/internal/db"
	"asset-service/internal/dtos"
	"asset-service/proto"

	"github.com/jackc/pgx/v5/pgtype"
)

// --- ASSET CLASS ---

func (s *AssetServer) CreateAssetClass(ctx context.Context, req *proto.CreateAssetClassRequest) (*proto.AssetClassResponse, error) {
	dto := dtos.CreateAssetClass{
		Name:       req.Name,
		Definition: req.Definition,
	}

	res, err := s.Queries.CreateAssetClass(ctx, db.CreateAssetClassParams{
		Name: dto.Name,
		Definition: pgtype.Text{String: dto.Definition, Valid: dto.Definition != ""},
	})
	if err != nil {
		return nil, err
	}

	return &proto.AssetClassResponse{Id: res.ID, Name: res.Name, Definition: res.Definition.String}, nil
}

func (s *AssetServer) UpdateAssetClass(ctx context.Context, req *proto.UpdateAssetClassRequest) (*proto.AssetClassResponse, error) {
	current, err := s.Queries.GetAssetClassByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	dto := dtos.UpdateAssetClass{
		Name:       &current.Name,
		Definition: &current.Definition.String,
	}

	if req.Name != nil { dto.Name = req.Name }
	if req.Definition != nil { dto.Definition = req.Definition }

	updated, err := s.Queries.UpdateAssetClass(ctx, db.UpdateAssetClassParams{
		ID:   current.ID,
		Name: *dto.Name,
		Definition: pgtype.Text{String: *dto.Definition, Valid: dto.Definition != nil},
	})
	if err != nil {
		return nil, err
	}

	return &proto.AssetClassResponse{Id: updated.ID, Name: updated.Name, Definition: updated.Definition.String}, nil
}

func (s *AssetServer) ListAssetClasses(ctx context.Context, req *proto.Empty) (*proto.ListAssetClassesResponse, error) {
	classes, err := s.Queries.ListAssetClasses(ctx)
	if err != nil {
		return nil, err
	}

	var res []*proto.AssetClassResponse
	for _, c := range classes {
		res = append(res, &proto.AssetClassResponse{Id: c.ID, Name: c.Name, Definition: c.Definition.String})
	}
	return &proto.ListAssetClassesResponse{Classes: res}, nil
}

// --- CRITICALITY ---

func (s *AssetServer) CreateCriticality(ctx context.Context, req *proto.CreateCriticalityRequest) (*proto.CriticalityResponse, error) {
	dto := dtos.CreateCriticality{
		Name:  req.Name,
		Value: int(req.Value),
	}

	res, err := s.Queries.CreateAssetCriticality(ctx, db.CreateAssetCriticalityParams{
		Name:  dto.Name,
		Value: int32(dto.Value),
	})
	if err != nil {
		return nil, err
	}

	return &proto.CriticalityResponse{Id: res.ID, Name: res.Name, Value: res.Value}, nil
}

func (s *AssetServer) UpdateCriticality(ctx context.Context, req *proto.UpdateCriticalityRequest) (*proto.CriticalityResponse, error) {
	current, err := s.Queries.GetAssetCriticalityByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	val := int(current.Value)
	dto := dtos.UpdateCriticality{
		Name:  &current.Name,
		Value: &val,
	}

	if req.Name != nil { dto.Name = req.Name }
	if req.Value != nil {
		v := int(*req.Value)
		dto.Value = &v
	}

	updated, err := s.Queries.UpdateAssetCriticality(ctx, db.UpdateAssetCriticalityParams{
		ID:    current.ID,
		Name:  *dto.Name,
		Value: int32(*dto.Value),
	})
	if err != nil {
		return nil, err
	}

	return &proto.CriticalityResponse{Id: updated.ID, Name: updated.Name, Value: updated.Value}, nil
}

// --- DEVICE TYPE ---

func (s *AssetServer) CreateDeviceType(ctx context.Context, req *proto.CreateDeviceTypeRequest) (*proto.DeviceTypeResponse, error) {
	dto := dtos.CreateDeviceType{
		AssetClassID: int(req.AssetClassId),
		Name:         req.Name,
	}

	res, err := s.Queries.CreateDeviceType(ctx, db.CreateDeviceTypeParams{
		AssetClassID: int32(dto.AssetClassID),
		Name:         dto.Name,
	})
	if err != nil {
		return nil, err
	}

	return &proto.DeviceTypeResponse{Id: res.ID, Name: res.Name, AssetClassId: res.AssetClassID}, nil
}

func (s *AssetServer) ListDeviceTypes(ctx context.Context, req *proto.Empty) (*proto.ListDeviceTypesResponse, error) {
	types, err := s.Queries.ListDeviceTypes(ctx)
	if err != nil {
		return nil, err
	}

	var res []*proto.DeviceTypeResponse
	for _, t := range types {
		res = append(res, &proto.DeviceTypeResponse{Id: t.ID, Name: t.Name, AssetClassId: t.AssetClassID})
	}
	return &proto.ListDeviceTypesResponse{DeviceTypes: res}, nil
}

// --- DELETE HANDLERS ---

func (s *AssetServer) DeleteAssetClass(ctx context.Context, req *proto.DeleteMetadataRequest) (*proto.Empty, error) {
	err := s.Queries.DeleteAssetClass(ctx, req.Id)
	return &proto.Empty{}, err
}

func (s *AssetServer) DeleteCriticality(ctx context.Context, req *proto.DeleteMetadataRequest) (*proto.Empty, error) {
	err := s.Queries.DeleteAssetCriticality(ctx, req.Id)
	return &proto.Empty{}, err
}
