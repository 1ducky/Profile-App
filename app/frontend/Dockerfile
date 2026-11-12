
# Build stage
FROM node:22-alpine AS builder

WORKDIR /app

# Enable pnpm through Corepack
RUN corepack enable

# Copy dependency manifests
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN pnpm install --frozen-lockfile

# Copy source
COPY . .

ARG API_URL
ENV API_URL=$API_URL
ARG APP_ORIGIN
ENV APP_ORIGIN=$APP_ORIGIN
# Build Astro SSR
RUN pnpm build


# Production stage
FROM node:22-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV HOST=0.0.0.0
ENV PORT=4321

# Copy build output and runtime dependencies
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

EXPOSE 4321

CMD ["node", "./dist/server/entry.mjs"]
