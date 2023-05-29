FROM golang:1.20-bullseye as builder
WORKDIR /go/src
COPY . /go/src/
RUN set -e \
    && export GOPROXY=https://goproxy.cn,direct \
    && go mod download \
    && go build -o linear cmd/main.go

FROM debian:bullseye
WORKDIR /app
COPY --from=builder /go/src/linear .
COPY ./configs /app/configs
RUN set -e \
    && apt update -yqq \
    && apt install -yqq ca-certificates \
    && apt clean autoclean \
    && apt autoremove -yqq \
    && rm -rf /var/lib/apt/lists/*
EXPOSE 8080
CMD ["/app/linear"]