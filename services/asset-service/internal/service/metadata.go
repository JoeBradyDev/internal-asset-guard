package service

import (
	"context"

	"asset-service/internal/db"
	"asset-service/proto"

	"github.com/jackc/pgx/v5/pgtype"
)

// --- ASSET CLASS ---

func (s *AssetServer) CreateAssetClass(ctx context.Context, req *proto.CreateAssetClassRequest) (*proto.AssetClassResponse, error) {
	res, err := s.Queries.CreateAssetClass(ctx, db.CreateAssetClassParams{
		Name: req.Name,
		Definition: pgtype.Text{String: req.Definition, Valid: req.Definition != ""},
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

	name := current.Name
	defn := current.Definition.String

	if req.Name != nil { name = *req.Name }
	if req.Definition != nil { defn = *req.Definition }

	updated, err := s.Queries.UpdateAssetClass(ctx, db.UpdateAssetClassParams{
		ID:         current.ID,
		Name:       name,
		Definition: pgtype.Text{String: defn, Valid: true},
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
	res, err := s.Queries.CreateAssetCriticality(ctx, db.CreateAssetCriticalityParams{
		Name:  req.Name,
		Value: int32(req.Value),
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

	name := current.Name
	val := current.Value

	if req.Name != nil { name = *req.Name }
	if req.Value != nil { val = int32(*req.Value) }

	updated, err := s.Queries.UpdateAssetCriticality(ctx, db.UpdateAssetCriticalityParams{
		ID:    current.ID,
		Name:  name,
		Value: val,
	})
	if err != nil {
		return nil, err
	}

	return &proto.CriticalityResponse{Id: updated.ID, Name: updated.Name, Value: updated.Value}, nil
}

// --- DEVICE TYPE ---

func (s *AssetServer) CreateDeviceType(ctx context.Context, req *proto.CreateDeviceTypeRequest) (*proto.DeviceTypeResponse, error) {
	res, err := s.Queries.CreateDeviceType(ctx, db.CreateDeviceTypeParams{
		AssetClassID: req.AssetClassId,
		Name:         req.Name,
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
