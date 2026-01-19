# Stage 1: Build
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npx nx build frontend --prod

# Stage 2: Runner
FROM node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production
ENV PORT 4200
# THIS IS THE CRITICAL LINE:
ENV HOSTNAME "0.0.0.0"

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/dist/apps/frontend/.next/standalone ./
COPY --from=builder /app/dist/apps/frontend/.next/static ./dist/apps/frontend/.next/static
COPY --from=builder /app/dist/apps/frontend/public ./dist/apps/frontend/public

USER nextjs
EXPOSE 4200

# Standalone creates a server.js that respects the HOSTNAME env var
CMD ["node", "apps/frontend/server.js"]
