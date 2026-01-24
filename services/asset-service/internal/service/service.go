package service

import (
	"asset-service/internal/db"
	"asset-service/proto"
)

// AssetServer is the implementation of the AssetService gRPC server.
// It embeds UnimplementedAssetServiceServer for forward compatibility
// and holds the SQLC Queries for database access.
type AssetServer struct {
	proto.UnimplementedAssetServiceServer
	Queries *db.Queries
}

// NewAssetServer creates a new instance of the service.
// This is called in your main.go to initialize the server.
func NewAssetServer(queries *db.Queries) *AssetServer {
	return &AssetServer{
		Queries: queries,
	}
}
