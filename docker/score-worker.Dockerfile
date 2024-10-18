FROM golang:1.23-alpine3.19 AS development
WORKDIR /go/src/main

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64
ENV GOCACHE=/root/.cache/go-build
ENV GOMODCACHE=/go/pkg/mod

RUN apk update && apk add --no-cache git
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
COPY . .

CMD [ "air", "-c", ".air.toml" ]