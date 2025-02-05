# build stage
FROM node:22.13 AS builder

ARG ORIGIN
RUN if [ -z "$ORIGIN" ]; then echo "ERROR: ORIGIN is not set!"; exit 1; fi

WORKDIR /app
COPY package*.json ./

RUN npm ci
COPY . .

RUN npm run build

# runtime stage
FROM node:22.13-alpine

WORKDIR /app
COPY --from=builder /app/build ./build
COPY package*.json ./

RUN npm ci --omit=dev

#RUN apt-get update && \
#    apt-get install -y --no-install-recommends curl && \
#    apt-get clean && \
#    rm -rf /var/lib/apt/lists/*

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://localhost:3000/health" ]

EXPOSE 3000
ENTRYPOINT [ "node", "./build" ]