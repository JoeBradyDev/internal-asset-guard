# Stage 1: Build (uses root context for Nx workspace access)
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npx nx build gateway --prod

# Stage 2: Production (Slim)
FROM node:20-alpine
WORKDIR /app
# Copy ONLY the compiled main.js and package.json
COPY --from=builder /app/dist/apps/gateway .
COPY package*.json ./
RUN npm install --omit=dev

EXPOSE 3000
CMD ["node", "main.js"]
