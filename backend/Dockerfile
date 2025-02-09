# build stage
FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o main ./cmd/server/main.go

# runtime stage
FROM alpine:3.21.2
WORKDIR /app
COPY --from=builder /build/main .

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8082
ENTRYPOINT [ "./main" ]
# CMD [""]
