# Build
FROM golang:1.21 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build make build

# Run
FROM ubuntu
WORKDIR /app
COPY --from=builder /src/server /app/server
ENTRYPOINT ["/app/server"]
