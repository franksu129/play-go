# Build Stage
FROM        golang as builder
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download && go build -o app ./cmd/gochat
#ENTRYPOINT  ["./app"]

# Final Stage
FROM gcr.io/distroless/base:latest
WORKDIR     /usr/local/bin

COPY --from=builder /app/app ./
COPY --from=builder /app/config /app/config
COPY --from=builder /app/web ./web

CMD ["/usr/local/bin/app"]