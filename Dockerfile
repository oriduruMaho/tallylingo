FROM golang:1-bullseye AS builder
WORKDIR /work
ARG CGO_ENABLED=0
COPY . .
RUN go build -o tallylingo ./cmd/main
FROM alpine:latest
ARG VERSION=0.5.1
LABEL org.opencontainers.image.source=https://github.com/oriduruMaho/tallylingo \
org.opencontainers.image.version=${VERSION} \
org.opencontainers.image.title=tallylingo 
# RUN adduser --disabled-password --disabled-login --home /workdir nonroot \
# && mkdir -p /workdir
RUN adduser -D -h /workdir nonroot && mkdir -p /workdir
COPY --from=builder /work/tallylingo /opt/tallylingo/tallylingo
COPY --from=golang:1.12 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /workdir
USER nonroot
ENTRYPOINT [ "/opt/tallylingo/tallylingo" ] 