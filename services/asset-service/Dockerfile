# Stage 1: Build
FROM golang:1.25-alpine3.21 AS builder

# Install protoc and Go gRPC plugins
RUN apk add --no-cache protoc protobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /app

# 1. Copy the shared proto library (requires root context in Tilt)
COPY libs/shared/proto/ ./libs/shared/proto/

# 2. Copy and download service dependencies
COPY services/asset-service/go.mod services/asset-service/go.sum ./services/asset-service/
WORKDIR /app/services/asset-service
RUN go mod download

# 3. Copy service source code
COPY services/asset-service/ .

# 4. Run the 'comment' command to generate gRPC code and build
RUN go generate ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /asset-service .

# Stage 2: Final
FROM alpine:3.21
WORKDIR /
COPY --from=builder /asset-service /asset-service

# Use standard gRPC port from apps.yaml
EXPOSE 50051
ENTRYPOINT ["/asset-service"]
