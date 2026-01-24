# Stage 1: Build
# Using a modern Go version (1.25) on Alpine 3.21
FROM golang:1.25-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /asset-service .

# Stage 2: Final
# Match the Alpine version exactly (3.21)
FROM alpine:3.21
WORKDIR /
COPY --from=builder /asset-service /asset-service

# Use a numeric ID for the user to stay shell-agnostic
USER 1000:1000
EXPOSE 50051
ENTRYPOINT ["/asset-service"]
