FROM golang:1.14 AS builder
WORKDIR /sky
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o /sky/bin/sky /sky/cmd/
FROM alpine:latest AS final
WORKDIR /app
COPY --from=builder /sky /app/
RUN cat /app/config/config.toml
RUN chmod +x /app/bin/sky
CMD ["/app/bin/sky","start","--path=/app/config/config.toml","--config=/app/build/config"]