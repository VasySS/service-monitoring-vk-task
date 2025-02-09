# build stage
FROM node:22.13-alpine AS builder

ARG ORIGIN
RUN if [ -z "$ORIGIN" ]; then echo "ERROR: ORIGIN is not set!"; exit 1; fi

ARG BACKEND_URL
RUN if [ -z "$BACKEND_URL" ]; then echo "ERROR: BACKEND_URL is not set!"; exit 1; fi

ENV VITE_BACKEND_URL=$BACKEND_URL

WORKDIR /app
COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build && npm prune --omit=dev

# runtime stage
FROM node:22.13-alpine

WORKDIR /app
COPY --from=builder --chown=node:node /app/build ./build
COPY --from=builder --chown=node:node /app/node_modules ./node_modules

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 3000
ENTRYPOINT [ "node", "build/index.js" ]