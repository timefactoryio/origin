# Build binary for target platform
FROM --platform=$BUILDPLATFORM golang:latest AS builder
WORKDIR /src
ARG TARGETARCH

# Download dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# Build static binary
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /out/origin .

# Copy CA certificates for HTTPS requests at runtime
RUN mkdir -p /out/etc/ssl/certs && \
    cp /etc/ssl/certs/ca-certificates.crt /out/etc/ssl/certs/

# Minimal scratch image — binary and certs only
FROM scratch
COPY --from=builder /out/origin /origin
COPY --from=builder /out/etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER 1000
ENTRYPOINT ["/origin"]