FROM golang:1.14 AS builder

WORKDIR /build
#RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY /home/wanglei/project/src/sky/go.sum .
RUN go mod download

COPY DockerFile .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o sky ./cmd/main.go

FROM alpine:3.10 AS final

WORKDIR /app
COPY --from=builder /build/sky /app/


# USER app-runner
ENTRYPOINT ["go run /app/cmd/main.go"]