# build stage
FROM golang:1.23.4 AS builder

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# runtime stage
FROM alpine:3.21.2
WORKDIR /app
COPY --from=builder /build/main .

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

ENTRYPOINT [ "./main" ]
