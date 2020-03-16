FROM golang:1.14 AS builder

WORKDIR /build

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o /build/bin/sky /build/cmd/main.go

FROM alpine:3.10 AS final

WORKDIR /app
COPY --from=builder /build /app/

ENTRYPOINT ["./bin/sky"]