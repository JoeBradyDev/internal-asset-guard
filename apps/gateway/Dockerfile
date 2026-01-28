# Stage 1: Build
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
# This build command generates dist/apps/gateway/asset/asset_service.proto
RUN npx nx build gateway --prod

# Stage 2: Production
FROM node:20-alpine
WORKDIR /app

# Copy the entire dist output (including the asset/ subfolder)
COPY --from=builder /app/dist/apps/gateway ./

# Copy package files and install production deps
COPY package*.json ./
RUN npm install --omit=dev

EXPOSE 3000
# main.js is at the root of /app, and protos are at /app/asset/
CMD ["node", "main.js"]
