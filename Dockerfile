# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /src

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source (templates are //go:embed-ed at build time)
COPY . .

# Build a fully static binary so it runs on a scratch/alpine base
ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /out/demo .

# ---- Runtime stage ----
FROM alpine:3.20

# A non-root user for runtime
RUN adduser -D -u 10001 appuser

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /out/demo /app/demo

# css/ is served via gin.StaticFile (read from disk at runtime), so it must exist.
# templates/ is embedded into the binary at build time, so it is NOT needed here.
COPY css ./css

RUN chown -R appuser:appuser /app
USER appuser

# The app reads the port from the "Port" env var (default 8080)
ENV Port=8080
EXPOSE 8080

ENTRYPOINT ["/app/demo"]
